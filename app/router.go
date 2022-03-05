package app

import (
	"github.com/accexs/github-microservice/controllers"
)

func initRouter() {
	server.POST("/top-secret", controllers.FormatMessage)
}
