package infrastructure

import (
	"api/src/application"
	"api/src/domain/repositories"
	"api/src/infrastructure/controller"
	rabbitService "api/src/infrastructure/services"
)

type Dependencies struct {
	ViewPlatillosController   controller.ViewPlatillosController
	CreatePlatilloController  controller.CreatePlatilloController
	UpdatePlatilloController  controller.UpdatePlatilloController
	DeletePlatilloController  controller.DeletePlatilloController
	NotificarPedidoController controller.NotificarPedidoController
}

func InitDependencies() Dependencies {
	repo := repositories.PlatilloRepository{}
	notificador := rabbitService.NewRabbitNotificador()

	return Dependencies{
		ViewPlatillosController: controller.ViewPlatillosController{
			UseCase: application.ViewPlatillosUseCase{Repo: repo},
		},
		CreatePlatilloController: controller.CreatePlatilloController{
			UseCase: application.CreatePlatilloUseCase{Repo: repo},
		},
		UpdatePlatilloController: controller.UpdatePlatilloController{
			UseCase: application.UpdatePlatilloUseCase{Repo: repo},
		},
		DeletePlatilloController: controller.DeletePlatilloController{
			UseCase: application.DeletePlatilloUseCase{Repo: repo},
		},
		NotificarPedidoController: controller.NotificarPedidoController{
			UseCase: application.NotificarPedidoCompletadoUseCase{
				Notificador: notificador,
			},
		},
	}
}
