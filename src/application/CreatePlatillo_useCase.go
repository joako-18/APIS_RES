package application

import (
	"api/src/domain/entities"
	"api/src/domain/repositories"
)

type CreatePlatilloUseCase struct {
	Repo repositories.PlatilloRepository
}

func (uc CreatePlatilloUseCase) Execute(p entities.Platillo) error {
	return uc.Repo.Create(p)
}
