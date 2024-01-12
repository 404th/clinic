package helper

import (
	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, s int, data interface{}) {
	if s >= 400 {
		c.AbortWithStatusJSON(s, data)
		return
	} else if s >= 200 && s < 300 {
		c.IndentedJSON(s, data)
		return
	}
}
