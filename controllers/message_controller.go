package controllers

import (
	"github.com/abbul/operacion-fuego-quasar/domain/top_secret"
	"github.com/abbul/operacion-fuego-quasar/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FormatTopSecret(ctx *gin.Context) {
	var request top_secret.Request
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	x, y, err := services.JoinLocations(request.Satellites)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	message, err := services.JoinMessages(request.Satellites)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"position": gin.H{"x": x, "y": y}, "message": message})
}

func AddInfoForTopSecretSplit(ctx *gin.Context) {
	var request top_secret.RequestSplit
	if err := ctx.ShouldBindUri(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid path request"})
		return
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid body request"})
		return
	}

	_, err := services.AddTopSecret(request)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"result": "OK"})
}

func FormatTopSecretSplit(ctx *gin.Context) {
	x, y, err := services.GetLocationSplit()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	message, err := services.GetMessageSplit()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"position": gin.H{"x": x, "y": y}, "message": message})
}
