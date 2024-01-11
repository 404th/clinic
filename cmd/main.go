package main

import (
	"fmt"

	"github.com/404th/clinic/api"
	"github.com/404th/clinic/api/handler"
	"github.com/404th/clinic/config"
	"github.com/404th/clinic/internal/service"
	"github.com/404th/clinic/internal/storage/postgres"
	log "github.com/sirupsen/logrus"
)

func main() {
	l := log.New()

	// ...1 => getting configurations
	cfg, err := config.GetConfig()
	if err != nil {
		l.Error(err)
		return
	}

	// ...2 => setting logger config
	l.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	if cfg.LogLevel == "debug" {
		l.SetLevel(log.DebugLevel)
	} else if cfg.LogLevel == "test" {
		l.SetLevel(log.DebugLevel)
	} else if cfg.LogLevel == "release" {
		l.SetLevel(log.InfoLevel)
	} else {
		l.SetLevel(log.DebugLevel)
	}

	// ...3 => connecting to psql
	strg, err := postgres.NewPostgres(cfg)
	if err != nil {
		l.Error(err)
		return
	}

	sv := service.NewService(cfg, l, strg)
	h := handler.NewHandler(cfg, l, sv)

	ginEngine := api.Run(cfg, h)
	if err := ginEngine.Run(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		l.Error(err)
		return
	}

}
