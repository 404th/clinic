package storage

import (
	"context"

	"github.com/404th/clinic/model"
)

type StorageI interface {
	UserStorage() UserI
}

type UserI interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error)
	Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error)
	GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error)
	UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.IDTracker, err error)
	DeleteUser(ctx context.Context, req *model.IDTracker) (err error)
	TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error)
}
