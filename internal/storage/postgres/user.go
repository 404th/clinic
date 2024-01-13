package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/util"
	"github.com/jackc/pgx/v4/pgxpool"
)

type user struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) storage.UserI {
	return &user{db}
}

func (o *user) CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	query := fmt.Sprintf(`
		INSERT INTO %s (
			role_id,
			username,
			firstname,
			surname,
			email,
			password
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) RETURNING id
	`, users_table_name)

	var id_sql sql.NullString

	if err = o.db.QueryRow(ctx, query,
		req.RoleID,
		req.Username,
		req.Firstname,
		req.Surname,
		req.Email,
		req.Password,
	).Scan(&id_sql); err != nil {
		return resp, err
	}

	if id_sql.Valid {
		resp.ID = id_sql.String
	}

	return resp, err
}

func (o *user) Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error) {
	resp = &model.LoginResponse{}

	var (
		filter = ` WHERE deleted_at IS NULL AND username=$1`
		query  = `SELECT id, password FROM "users" `
	)

	query += filter

	var (
		id_sql       sql.NullString
		password_sql sql.NullString
		id           string
		password     string
	)

	if err = o.db.QueryRow(ctx, query, req.Username).Scan(&id_sql, &password_sql); err != nil {
		if strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
			return resp, nil
		}

		return resp, err
	}

	if id_sql.Valid {
		id = id_sql.String
	}

	if password_sql.Valid {
		password = password_sql.String
	}

	if !util.CheckPasswordHash(req.Password, password) {
		return resp, errors.New("invalid password")
	}

	resp.ID = id

	return resp, err
}

func (o *user) GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error) {
	resp = &model.User{}

	query := fmt.Sprintf(`
		SELECT 
			id,
			role_id,
			username,
			firstname,
			surname,
			email,
			wallet,
			active
		FROM 
			%s 
		WHERE deleted_at IS NULL AND id = $1
	`, users_table_name)

	var usr model.User
	var (
		firstname_sql sql.NullString
		surname_sql   sql.NullString
		email_sql     sql.NullString
		wallet_sql    sql.NullFloat64
		active_sql    sql.NullInt16
	)

	if err = o.db.QueryRow(ctx, query, req.ID).Scan(
		&usr.ID,
		&usr.RoleID,
		&usr.Username,
		&firstname_sql,
		&surname_sql,
		&email_sql,
		&wallet_sql,
		&active_sql,
	); err != nil {
		if strings.Contains(sql.ErrNoRows.Error(), err.Error()) {
			return resp, errors.New("user not found")
		}
		return resp, err
	}

	if firstname_sql.Valid {
		usr.Firstname = firstname_sql.String
	}

	if surname_sql.Valid {
		usr.Surname = surname_sql.String
	}

	if firstname_sql.Valid {
		usr.Firstname = firstname_sql.String
	}

	if email_sql.Valid {
		usr.Email = email_sql.String
	}

	if wallet_sql.Valid {
		usr.Wallet = wallet_sql.Float64
	}

	if active_sql.Valid {
		usr.Active = int(active_sql.Int16)
	}

	resp = &usr

	return resp, nil
}

func (o *user) GetAllUsers(ctx context.Context, req *model.GetAllUsersRequest) (resp *model.GetAllUsersResponse, err error) {
	resp = &model.GetAllUsersResponse{}

	var (
		offset int32
	)

	offset = (req.Page - 1) * req.Limit

	query := fmt.Sprintf(`
		SELECT 
			id,
			role_id,
			username,
			firstname,
			surname,
			email,
			wallet,
			active  
		FROM 
			%s 
		WHERE deleted_at IS NULL 
		ORDER BY updated_at DESC
		LIMIT $1 OFFSET $2 
	`, users_table_name)

	rows, err := o.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		var (
			firstname_sql sql.NullString
			surname_sql   sql.NullString
			email_sql     sql.NullString
			wallet_sql    sql.NullFloat64
			active_sql    sql.NullInt16
		)

		if err = rows.Scan(
			&user.ID,
			&user.RoleID,
			&user.Username,
			&firstname_sql,
			&surname_sql,
			&email_sql,
			&wallet_sql,
			&active_sql,
		); err != nil {
			return resp, err
		}

		if firstname_sql.Valid {
			user.Firstname = firstname_sql.String
		}

		if surname_sql.Valid {
			user.Surname = surname_sql.String
		}

		if email_sql.Valid {
			user.Email = email_sql.String
		}

		if wallet_sql.Valid {
			user.Wallet = wallet_sql.Float64
		}

		if active_sql.Valid {
			user.Active = int(active_sql.Int16)
		}

		resp.Users = append(resp.Users, user)
	}

	count_query := fmt.Sprintf(`
		SELECT COUNT(*) as count FROM %s WHERE deleted_at IS NULL 
	`, users_table_name)

	var count int

	if err = o.db.QueryRow(ctx, count_query).Scan(&count); err != nil {
		return resp, err
	}

	resp.Metadata.Count = count
	resp.Metadata.Limit = req.Limit
	resp.Metadata.Page = req.Page

	return resp, err
}

func (o *user) TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET wallet = wallet + $1 
		WHERE id = $2
	`, users_table_name)

	r, err := o.db.Exec(ctx, query, req.Value, req.ID)
	if err != nil {
		return resp, err
	}

	if r.RowsAffected() <= 0 {
		return resp, errors.New("not updated")
	}

	resp.ID = req.ID

	return resp, err
}
