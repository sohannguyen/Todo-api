package controller

import "github.com/gin-gonic/gin"

func loginFunc(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func test() {

}
