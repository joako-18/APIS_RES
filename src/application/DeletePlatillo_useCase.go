package application

import "api/src/domain/repositories"

type DeletePlatilloUseCase struct {
	Repo repositories.PlatilloRepository
}

func (uc DeletePlatilloUseCase) Execute(id int) error {
	return uc.Repo.Delete(id)
}
