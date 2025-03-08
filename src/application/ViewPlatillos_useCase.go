package application

import (
	"api/src/domain/entities"
	"api/src/domain/repositories"
)

type ViewPlatillosUseCase struct {
	Repo repositories.PlatilloRepository
}

func (uc ViewPlatillosUseCase) Execute() ([]entities.Platillo, error) {
	return uc.Repo.GetAll()
}
