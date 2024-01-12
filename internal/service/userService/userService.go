package userService

import (
	"context"

	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/service"
	"github.com/404th/clinic/internal/storage"
	"github.com/404th/clinic/model"
	"github.com/sirupsen/logrus"
)

type userService struct {
	cfg  *config.Config
	log  *logrus.Logger
	strg storage.UserI
}

func NewUserService(cfg *config.Config, log *logrus.Logger, strg storage.UserI) service.UserServiceI {
	return &userService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (us *userService) CreateUser(ctx context.Context, req *model.CreateUserRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	us.log.Infof("CreateUser() => username: %s => req: %+v", req.Username, req)
	resp, err = us.strg.CreateUser(ctx, req)
	if err != nil {
		us.log.Errorf("CreateUser() => username: %s => err: %+v", req.Username, err)
		return resp, err
	}

	us.log.Infof("CreateUser() => username: %s => resp: %+v", req.Username, req)
	return resp, err
}

func (us *userService) Login(ctx context.Context, req *model.LoginRequest) (resp *model.LoginResponse, err error) {
	resp = &model.LoginResponse{}

	return resp, err
}

func (us *userService) GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error) {
	return resp, err
}

func (us *userService) UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (resp *model.IDTracker, err error) {
	return resp, err
}

func (us *userService) DeleteUser(ctx context.Context, req *model.IDTracker) (err error) {
	return err
}

func (us *userService) TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error) {
	return resp, err
}
