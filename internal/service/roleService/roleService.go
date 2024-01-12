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
