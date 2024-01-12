package service

import (
	"context"

	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type service struct {
	cfg  *config.Config
	log  *logrus.Logger
	strg storage.StorageI
	db   *pgxpool.Pool
}

func NewService(cfg *config.Config, log *logrus.Logger, strg storage.StorageI) ServiceI {
	return &service{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

type ServiceI interface {
	UserService() UserServiceI
	RoleService() RoleServiceI
}

type UserServiceI interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error)
	Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error)
	GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error)
	UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.IDTracker, err error)
	DeleteUser(ctx context.Context, req *model.IDTracker) (err error)
	TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error)
}

type RoleServiceI interface {
	CreateRole(ctx context.Context, req *model.CreateRoleRequest) (resp *model.IDTracker, err error)
}

func (sc *service) UserService() UserServiceI {
	return sc.strg.UserStorage()
}

func (sc *service) RoleService() RoleServiceI {
	return sc.strg.RoleStorage()
}
