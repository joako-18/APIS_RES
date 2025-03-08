package controller

import (
	"api/src/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewPlatillosController struct {
	UseCase application.ViewPlatillosUseCase
}

func (c ViewPlatillosController) GetAll(ctx *gin.Context) {
	platillos, err := c.UseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, platillos)
}
