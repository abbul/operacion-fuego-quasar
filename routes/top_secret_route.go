package routes

import (
	"github.com/abbul/operacion-fuego-quasar/controllers"
	"github.com/gin-gonic/gin"
)

func AddTopSecretRoutes(rg *gin.RouterGroup) {
	rgTopSecretSplit := rg.Group("/top_secret")
	rgTopSecretSplit.POST("/", controllers.FormatTopSecret)
}
