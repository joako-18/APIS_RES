package application

import (
	"api/src/domain/entities"
	"api/src/domain/repositories"
)

type UpdatePlatilloUseCase struct {
	Repo repositories.PlatilloRepository
}

func (uc UpdatePlatilloUseCase) Execute(p entities.Platillo) error {
	return uc.Repo.Update(p)
}
