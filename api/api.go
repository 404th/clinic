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
	r.POST("/refresh-token", h.RefreshToken)

	user := r.Group("/user", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		user.GET("/:id", h.GetUserByID)
		user.PATCH("/", h.TransferMoney)
		user.PUT("/transfer", h.TransferMoney)
	}

	role := r.Group("/role", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		role.GET("/:id")
	}

	queue := r.Group("/queue", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		queue.POST("/", h.CreateQueue)
		queue.GET("/:id")
		queue.PATCH("/", h.MakePurchase)
	}

	return r
}
