package api

import (
	"github.com/404th/clinic/api/handler"
	"github.com/404th/clinic/api/middleware"
	"github.com/404th/clinic/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config, h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/:id", h.CreateUser)
	r.POST("/user", h.Login)
	usr := r.Group("/user", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		usr.GET("/:id", h.GetUserByID)
		usr.PUT("/", h.UpdateUser)
		usr.PATCH("/", h.TransferMoney)
		usr.DELETE("/", h.DeleteUser)
	}

	return r
}
