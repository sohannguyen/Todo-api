package routers

import (
	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

var router = gin.Default()

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	// Auth router
	auth := r.router.Group("/auth")
	r.authRouter(auth)

	// User router
	user := r.router.Group("/user")
	r.userRouter(user)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}
