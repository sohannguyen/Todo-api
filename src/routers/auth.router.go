package routers

import (
	"github.com/gin-gonic/gin"
)

func (r routes) authRouter(rg *gin.RouterGroup) {
	rg.POST("/login")
}

func pongFunction(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
