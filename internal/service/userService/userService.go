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

	us.log.Infof("Login() => username: %s => req: %+v", req.Username, req)
	resp, err = us.strg.Login(ctx, req)
	if err != nil {
		us.log.Errorf("Login() => username: %s => err: %+v", req.Username, err)
		return resp, err
	}

	us.log.Infof("Login() => username: %s => resp: %+v", req.Username, req)
	return resp, err
}

func (us *userService) GetUserByID(ctx context.Context, req *model.IDTracker) (resp *model.User, err error) {
	resp = &model.User{}

	us.log.Infof("GetUserByID() => id: %s => req: %+v", req.ID, req)
	resp, err = us.strg.GetUserByID(ctx, req)
	if err != nil {
		us.log.Errorf("GetUserByID() => id: %s => err: %+v", req.ID, err)
		return resp, err
	}

	us.log.Infof("GetUserByID() => username: %s => resp: %+v", resp.Username, resp)
	return resp, err
}

func (us *userService) GetAllUsers(ctx context.Context, req *model.GetAllUsersRequest) (resp *model.GetAllUsersResponse, err error) {
	resp = &model.GetAllUsersResponse{}

	us.log.Infof("GetAllUsers() => page: %d => req: %+v", req.Page, req)
	resp, err = us.strg.GetAllUsers(ctx, req)
	if err != nil {
		us.log.Errorf("GetAllUsers() => page: %d => err: %+v", req.Page, err)
		return resp, err
	}

	us.log.Infof("GetAllUsers() => metadata: %+v => resp: %+v", resp.Metadata, resp)
	return resp, err
}

func (us *userService) TransferMoney(ctx context.Context, req *model.TransferMoneyRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	us.log.Infof("TransferMoney() => id: %s => req: %+v", req.ID, req)
	resp, err = us.strg.TransferMoney(ctx, req)
	if err != nil {
		us.log.Errorf("TransferMoney() => id: %s => err: %+v", req.ID, err)
		return resp, err
	}

	us.log.Infof("TransferMoney() => id: %s => resp: %+v", resp.ID, resp)
	return resp, err
}
