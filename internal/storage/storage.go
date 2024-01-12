package storage

import (
	"context"

	"github.com/404th/clinic/model"
)

type StorageI interface {
	UserStorage() UserI
	RoleStorage() RoleI
	QueueStorage() QueueI
}

type UserI interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error)
	Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error)
	GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error)
	TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error)
}

type RoleI interface {
	CreateRole(ctx context.Context, req *model.CreateRoleRequest) (resp *model.IDTracker, err error)
}

type QueueI interface {
	CreateQueue(ctx context.Context, req *model.CreateQueueRequest) (resp *model.IDTracker, err error)
	MakePurchase(ctx context.Context, req *model.MakePurchaseRequest) (resp *model.IDTracker, err error)
}
