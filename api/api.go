package api

import (
	"github.com/404th/clinic/api/handler"
	"github.com/404th/clinic/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config, h *handler.Handler) *gin.Engine {
	r := gin.Default()

	usr := r.Group("/user")
	{
		usr.POST("/")
		usr.GET("/:id")
		usr.PUT("/")
		usr.PATCH("/")
		usr.DELETE("/")
	}

	return r
}
