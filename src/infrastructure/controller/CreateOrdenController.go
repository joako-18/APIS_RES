package controller

import (
	"api/src/infrastructure"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrdenController struct {
	RabbitMQ *infrastructure.RabbitMQService
}

func (c OrdenController) Ordenar(ctx *gin.Context) {
	var request struct {
		PlatilloID int `json:"platillo_id"`
	}

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := fmt.Sprintf(`{"platillo_id": %d}`, request.PlatilloID)
	if err := c.RabbitMQ.PublishMessage(message); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error enviando orden a RabbitMQ"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Orden enviada a RabbitMQ"})
}
