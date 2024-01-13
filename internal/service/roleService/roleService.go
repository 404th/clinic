package roleService

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
	strg storage.RoleI
}

func NewUserService(cfg *config.Config, log *logrus.Logger, strg storage.RoleI) service.RoleServiceI {
	return &userService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (us *userService) CreateRole(ctx context.Context, req *model.CreateRoleRequest) (resp *model.IDTracker, err error) {
	resp = &model.IDTracker{}

	us.log.Infof("CreateRole() => rolename: %s => req: %+v", req.Rolename, req)
	resp, err = us.strg.CreateRole(ctx, req)
	if err != nil {
		us.log.Errorf("CreateRole() => rolename: %s => err: %+v", req.Rolename, err)
		return resp, err
	}

	us.log.Infof("CreateRole() => rolename: %s => resp: %+v", req.Rolename, req)
	return resp, err
}

func (us *userService) GetAllRoles(ctx context.Context, req *model.GetAllRolesRequest) (resp *model.GetAllRolesResponse, err error) {
	resp = &model.GetAllRolesResponse{}

	us.log.Infof("GetAllRoles() => page: %d => req: %+v", req.Page, req)
	resp, err = us.strg.GetAllRoles(ctx, req)
	if err != nil {
		us.log.Errorf("GetAllRoles() => page: %d => err: %+v", req.Page, err)
		return resp, err
	}

	us.log.Infof("GetAllRoles() => metadata: %+v => resp: %+v", resp.Metadata, resp)
	return resp, err
}
