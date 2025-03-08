package infrastructure

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, deps Dependencies) {
	r.GET("/platillos", deps.ViewPlatillosController.GetAll)
	r.POST("/platillos", deps.CreatePlatilloController.Create)
	r.PUT("/platillos/:id", deps.UpdatePlatilloController.Update)
	r.DELETE("/platillos/:id", deps.DeletePlatilloController.Delete)
	r.POST("/ordenar", deps.OrdenController.Ordenar)
}
