package controller

import (
	"api/src/application"
	"api/src/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreatePlatilloController struct {
	UseCase application.CreatePlatilloUseCase
}

func (c CreatePlatilloController) Create(ctx *gin.Context) {
	var platillo entities.Platillo
	if err := ctx.ShouldBindJSON(&platillo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UseCase.Execute(platillo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Platillo creado correctamente"})
}
