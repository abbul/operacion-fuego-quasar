package app

import (
	"github.com/abbul/operacion-fuego-quasar/routes"
	"github.com/gin-gonic/gin"
)

var (
	engine *gin.Engine
)

func StartApp() (engine *gin.Engine) {
	engine = gin.Default()
	var prefixPath = ""
	routes.AddTopSecretRoutes(engine.Group(prefixPath))
	routes.AddTopSecretSplitRoutes(engine.Group(prefixPath))
	return engine
}
