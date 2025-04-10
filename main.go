package main

import (
	"api/src/core"
	"api/src/infrastructure"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	core.InitDB()

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	deps := infrastructure.InitDependencies()
	infrastructure.RegisterRoutes(r, deps)

	fmt.Println("Servidor corriendo en el puerto 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error iniciando el servidor:", err)
	}
}
