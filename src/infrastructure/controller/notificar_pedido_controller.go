package controller

import (
	"api/src/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotificarPedidoController struct {
	UseCase application.NotificarPedidoCompletadoUseCase
}

func (c *NotificarPedidoController) Notify(ctx *gin.Context) {
	var payload struct {
		PedidoID int    `json:"pedido_id"`
		Estado   string `json:"estado"`
	}

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	err := c.UseCase.Execute(payload.PedidoID, payload.Estado)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo notificar"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Notificación enviada"})
}
