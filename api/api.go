package api

import (
	"github.com/404th/clinic/api/handler"
	"github.com/404th/clinic/api/middleware"
	"github.com/404th/clinic/config"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config, h *handler.Handler) *gin.Engine {
	r := gin.Default()

	r.POST("/user", h.CreateUser)
	r.POST("/role", h.CreateRole)
	r.POST("/login", h.Login)
	usr := r.Group("/user", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		usr.GET("/:id", h.GetUserByID)
		usr.PUT("/", h.UpdateUser)
		usr.PATCH("/", h.TransferMoney)
		usr.DELETE("/", h.DeleteUser)
	}

	rl := r.Group("/role", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		rl.GET("/:id")
		rl.PUT("/")
		rl.PATCH("/")
		rl.DELETE("/")
	}

	return r
}
