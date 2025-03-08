package infrastructure

import (
	"api/src/application"
	"api/src/domain/repositories"
	"api/src/infrastructure/controller"
	"log"
)

type Dependencies struct {
	ViewPlatillosController  controller.ViewPlatillosController
	CreatePlatilloController controller.CreatePlatilloController
	UpdatePlatilloController controller.UpdatePlatilloController
	DeletePlatilloController controller.DeletePlatilloController
	OrdenController          controller.OrdenController
}

func InitDependencies() Dependencies {
	repo := repositories.PlatilloRepository{}

	rabbitMQ, err := rabbitmq.NewRabbitMQService()
	if err != nil {
		log.Fatal("Error inicializando RabbitMQ:", err)
	}

	return Dependencies{
		ViewPlatillosController:  controller.ViewPlatillosController{UseCase: application.ViewPlatillosUseCase{Repo: repo}},
		CreatePlatilloController: controller.CreatePlatilloController{UseCase: application.CreatePlatilloUseCase{Repo: repo}},
		UpdatePlatilloController: controller.UpdatePlatilloController{UseCase: application.UpdatePlatilloUseCase{Repo: repo}},
		DeletePlatilloController: controller.DeletePlatilloController{UseCase: application.DeletePlatilloUseCase{Repo: repo}},
		OrdenController:          controller.OrdenController{RabbitMQ: rabbitMQ},
	}
}
