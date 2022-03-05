package app

import (
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	server = gin.Default()
}

func StartApp() {
	initRouter()

	if err := server.Run(":8080"); err != nil {
		panic(err)
	}
}
