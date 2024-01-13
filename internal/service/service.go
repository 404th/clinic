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
	QueueService() QueueServiceI
}

type UserServiceI interface {
	CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error)
	Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error)
	GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error)
	GetAllUsers(ctx context.Context, req *model.GetAllUsersRequest) (resp *model.GetAllUsersResponse, err error)
	TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error)
}

type RoleServiceI interface {
	CreateRole(ctx context.Context, req *model.CreateRoleRequest) (resp *model.IDTracker, err error)
	GetAllRoles(ctx context.Context, req *model.GetAllRolesRequest) (resp *model.GetAllRolesResponse, err error)
}

type QueueServiceI interface {
	CreateQueue(ctx context.Context, req *model.CreateQueueRequest) (resp *model.IDTracker, err error)
	MakePurchase(ctx context.Context, req *model.MakePurchaseRequest) (resp *model.IDTracker, err error)
	GetAllQueues(ctx context.Context, req *model.GetAllQueuesRequest) (resp *model.GetAllQueuesResponse, err error)
}

func (sc *service) UserService() UserServiceI {
	return sc.strg.UserStorage()
}

func (sc *service) RoleService() RoleServiceI {
	return sc.strg.RoleStorage()
}

func (sc *service) QueueService() QueueServiceI {
	return sc.strg.QueueStorage()
}
