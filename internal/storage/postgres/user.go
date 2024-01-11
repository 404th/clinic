package postgres

import (
	"context"

	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type user struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) storage.UserI {
	return &user{db}
}

func (o *user) CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error) {
	return resp, err
}

func (o *user) Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error) {
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
