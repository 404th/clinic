package api

import (
	_ "github.com/404th/clinic/api/docs"
	"github.com/404th/clinic/api/handler"
	"github.com/404th/clinic/api/middleware"
	"github.com/404th/clinic/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetUpRouter godoc
// @description				This is a api gateway for clinic
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
func Run(cfg *config.Config, h *handler.Handler) *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	r.Use(cors.New(config))
	r.Use(MaxAllowed(100))

	r.POST("/user", h.CreateUser)
	r.POST("/role", h.CreateRole)
	r.POST("/login", h.Login)
	r.POST("/refresh-token", h.RefreshToken)

	user := r.Group("/user", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		user.GET("/:id", h.GetUserByID)
		user.GET("/", h.GetAllUsers)
		user.PATCH("/", h.TransferMoney)
		user.PUT("/transfer", h.TransferMoney)
	}

	role := r.Group("/role", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		role.GET("/", h.GetAllRoles)
	}

	queue := r.Group("/queue", middleware.JwtAuthMiddleware(cfg.AccessTokenSecret))
	{
		queue.POST("/", h.CreateQueue)
		queue.PATCH("/", h.MakePurchase)
		queue.GET("/", h.GetAllQueues)
	}

	// swagger
	url := ginSwagger.URL("/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()
	}
}
