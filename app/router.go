package app

import (
	"github.com/abbul/operacion-fuego-quasar/controllers"
)

func initRouter() {
	server.POST("/top-secret", controllers.FormatMessage)
}
