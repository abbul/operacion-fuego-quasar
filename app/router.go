package app

import (
	"github.com/abbul/operacion-fuego-quasar/controllers"
)

func initRouter() {
	server.POST("/top_secret", controllers.FormatTopSecret)
	server.POST("/top_secret_split/:satellite_name", controllers.AddInfoForTopSecretSplit)
	server.GET("/top_secret_split", controllers.FormatTopSecretSplit)
}
