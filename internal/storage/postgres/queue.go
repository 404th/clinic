package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type queue struct {
	db *pgxpool.Pool
}

func NewQueue(db *pgxpool.Pool) storage.QueueI {
	return &queue{
		db: db,
	}
}

func (q *queue) CreateQueue(ctx context.Context, req *model.CreateQueueRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	query := fmt.Sprintf(`
		INSERT INTO %s (
			recipient_id,
			customer_id
		) VALUES (
			$1,
			$2
		) RETURNING id
	`, queues_table_name)

	var id string
	if q.db.QueryRow(ctx, query, req.RecipientID, req.CustomerID).Scan(&id); err != nil {
		return resp, err
	}

	resp.ID = id

	return resp, err
}

func (q *queue) MakePurchase(ctx context.Context, req *model.MakePurchaseRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	tx, err := q.db.Begin(ctx)
	if err != nil {
		return resp, err
	}

	q1 := `
		UPDATE users 
		SET wallet = wallet - $1 
		FROM queues 
		WHERE 
			users.deleted_at IS NULL AND 
			wallet > $1 AND 
			queues.id = $2 AND 
			queues.customer_id = users.id AND 
			payment_status = 0
	`

	r1, err := tx.Exec(ctx, q1, req.Amount, req.QueueID)
	if err != nil {
		return resp, err
	}
	if r1.RowsAffected() <= 0 {
		_ = tx.Rollback(ctx)
		return resp, errors.New("payment not accepted")
	}

	q2 := `
		UPDATE users  
		SET wallet = wallet + $1 
		FROM queues 
		WHERE 
			queues.recipient_id = users.id AND 
			users.deleted_at IS NULL AND
			queues.deleted_at IS NULL AND 
			queues.id = $2 AND 
			payment_status = 0
	`

	r2, err := tx.Exec(ctx, q2, req.Amount, req.QueueID)
	if err != nil {
		return resp, err
	}
	if r2.RowsAffected() <= 0 {
		_ = tx.Rollback(ctx)
		return resp, errors.New("payment not accepted")
	}

	q3 := `
		UPDATE queues  
		SET paid_money = paid_money + $1
		WHERE queues.deleted_at IS NULL AND queues.id = $2 AND 
		payment_status = 0
	`

	r3, err := tx.Exec(ctx, q3, req.Amount, req.QueueID)
	if err != nil {
		return resp, err
	}
	if r3.RowsAffected() <= 0 {
		_ = tx.Rollback(ctx)
		return resp, errors.New("payment not accepted")
	}

	q4 := `
		UPDATE queues 
		SET 
			payment_status = CASE 
								WHEN paid_money >= subquery.role_price THEN 1
								ELSE 0
							END,
			queue_number = CASE 
							WHEN paid_money >= subquery.role_price THEN subquery.max_queue_number + 1 
							ELSE 0
						END
		FROM (
			SELECT 
				queues.recipient_id,
				MAX(queues.queue_number) as max_queue_number,
				roles.price as role_price
			FROM queues
			JOIN users ON users.id = queues.recipient_id
			JOIN roles ON roles.id = users.role_id
			WHERE queues.paid_money >= roles.price
			GROUP BY queues.recipient_id, roles.price
		) AS subquery
		WHERE 
			queues.recipient_id = subquery.recipient_id 
			AND queues.id = $1 
			AND queues.deleted_at IS NULL 
			AND payment_status = 0	
	`

	r4, err := tx.Exec(ctx, q4, req.QueueID)
	if err != nil {
		return resp, err
	}
	if r4.RowsAffected() <= 0 {
		_ = tx.Rollback(ctx)
		return resp, errors.New("payment not accepted")
	}

	resp.ID = req.QueueID

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		return resp, err
	}

	return resp, err
}

func (q *queue) GetAllQueues(ctx context.Context, req *model.GetAllQueuesRequest) (resp *model.GetAllQueuesResponse, err error) {
	resp = &model.GetAllQueuesResponse{}

	var (
		offset int32
	)

	offset = (req.Page - 1) * req.Limit

	query := fmt.Sprintf(`
		SELECT 
			id,
			recipient_id,
			customer_id,
			paid_money,
			queue_number,
			payment_status 
		FROM 
			%s 
		WHERE deleted_at IS NULL 
		ORDER BY updated_at DESC 
		LIMIT $1 OFFSET $2 
	`, queues_table_name)

	rows, err := q.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var queue model.Queue
		var (
			paid_money_sql     sql.NullFloat64
			queue_number_sql   sql.NullInt32
			payment_status_sql sql.NullInt32
		)

		if err = rows.Scan(
			&queue.ID,
			&queue.RecipientID,
			&queue.CustomerID,
			&paid_money_sql,
			&queue_number_sql,
			&payment_status_sql,
		); err != nil {
			return resp, err
		}

		if paid_money_sql.Valid {
			queue.PaidMoney = paid_money_sql.Float64
		}

		if queue_number_sql.Valid {
			queue.QueueNumber = int(queue_number_sql.Int32)
		}

		if payment_status_sql.Valid {
			queue.PaymentStatus = int(payment_status_sql.Int32)
		}

		resp.Queues = append(resp.Queues, queue)
	}

	count_query := fmt.Sprintf(`
		SELECT COUNT(*) as count FROM %s WHERE deleted_at IS NULL 
	`, queues_table_name)

	var count int

	if err = q.db.QueryRow(ctx, count_query).Scan(&count); err != nil {
		return resp, err
	}

	resp.Metadata.Count = count
	resp.Metadata.Limit = req.Limit
	resp.Metadata.Page = req.Page

	return resp, err
}
