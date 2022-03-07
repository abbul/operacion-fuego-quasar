package controllers

import (
	"github.com/abbul/operacion-fuego-quasar/domain/top_secret"
	"github.com/abbul/operacion-fuego-quasar/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormatMessage(ctx *gin.Context) {
	var request top_secret.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var x, y = services.JoinLocations(request.Satellites)
	var message = services.JoinMessages(request.Satellites)

	ctx.JSON(http.StatusOK, gin.H{"position": gin.H{"x": x, "y": y}, "message": message})
}
