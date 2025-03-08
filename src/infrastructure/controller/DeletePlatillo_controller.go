package controller

import (
	"api/src/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeletePlatilloController struct {
	UseCase application.DeletePlatilloUseCase
}

func (c DeletePlatilloController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.UseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Platillo eliminado"})
}
