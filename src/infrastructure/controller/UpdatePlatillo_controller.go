package controller

import (
	"api/src/application"
	"api/src/domain/entities"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdatePlatilloController struct {
	UseCase application.UpdatePlatilloUseCase
}

func (c UpdatePlatilloController) Update(ctx *gin.Context) {
	var platillo entities.Platillo
	if err := ctx.ShouldBindJSON(&platillo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	platillo.ID = id

	if err := c.UseCase.Execute(platillo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Platillo actualizado"})
}
