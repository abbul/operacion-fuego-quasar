package routes

import (
	"github.com/abbul/operacion-fuego-quasar/controllers"
	"github.com/gin-gonic/gin"
)

func AddTopSecretSplitRoutes(rg *gin.RouterGroup) {
	rgTopSecretSplit := rg.Group("/top_secret_split")
	rgTopSecretSplit.GET("/", controllers.FormatTopSecretSplit)
	rgTopSecretSplit.POST("/:satellite_name", controllers.AddInfoForTopSecretSplit)
}
