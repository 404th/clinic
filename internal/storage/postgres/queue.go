package postgres

import (
	"context"
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

var queues_table_name string = "queues"

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

	// "users" table have id role_id, wallet.
	// "roles" table have id, price.
	// "queues" table have id, paid_money, recipient_id, customer_id, payment_status.
	// (queues.recipient_id and queues.customer_id are related to users.id),
	// (users.role_id is related to roles.id)

	// two arguments are given from me: queue.id and value.
	// you should write query:
	// 1.  distract value from user's wallet which is related to the queue's customer_id with id if value is not higher than value of wallet field.
	// 2. add value to user's wallet which is related to the queue's recipient_id with id if value is not higher than user's wallet which is related to the queue's customer_id with id.
	// 3. add value to queue's paid_money field if value is not higher than user's wallet which is related to the queue's customer_id with id.
	// 4. after addition to paid_money, if queues' paid_money is higher or equal to role of recipient's price, then make queues' 'payment_status' field 1 else do nothing and if value is not higher than user's wallet which is related to the queue's customer_id with id.
	q1 := `
		
	  `

	r1, err := tx.Exec(ctx, q1, req.QueueID, req.Amount)
	if err != nil {
		return resp, err
	}
	if r1.RowsAffected() <= 0 {
		_ = tx.Rollback(ctx)
		return resp, errors.New("payment not accepted")
	}

	resp.ID = req.QueueID

	tx.Commit(ctx)
	return resp, err
}
