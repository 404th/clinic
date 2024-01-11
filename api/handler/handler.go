package handler

import (
	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/service"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	cfg     *config.Config
	log     *logrus.Logger
	service service.ServiceI
}

func NewHandler(cfg *config.Config, log *logrus.Logger, service service.ServiceI) *Handler {
	return &Handler{
		cfg:     cfg,
		log:     log,
		service: service,
	}
}
