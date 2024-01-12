package middleware

import (
	"net/http"
	"strings"

	"github.com/404th/clinic/internal/jwt"
	"github.com/404th/clinic/model"
	"github.com/404th/clinic/pkg/helper"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := jwt.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := jwt.ExtractIDFromToken(authToken, secret)
				if err != nil {
					helper.SendResponse(c, http.StatusUnauthorized, model.ErrorResponse{
						Message: err.Error(),
						Data:    err,
					})
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			helper.SendResponse(c, http.StatusUnauthorized, model.ErrorResponse{
				Message: err.Error(),
				Data:    err,
			})
			return
		}
		helper.SendResponse(c, http.StatusUnauthorized, model.ErrorResponse{
			Message: "User unauthorized",
			Data:    nil,
		})
	}
}
