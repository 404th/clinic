package postgres

import (
	"context"
	"database/sql"
	"errors"

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

	query := `
		INSERT INTO "users" (
			role_id,
			username,
			firstname,
			surname,
			email,
			password
		) VALUE (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6
		) RETURNING id
	`

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
	return resp, err
}

func (o *user) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.IDTracker, err error) {
	return resp, err
}

func (o *user) DeleteUser(ctx context.Context, req *model.IDTracker) (err error) {
	return err
}

func (o *user) TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error) {
	return resp, err
}
