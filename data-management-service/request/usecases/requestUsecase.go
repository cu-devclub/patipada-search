package usecases

import (
	"data-management/request/repositories"
	validator "data-management/structValidator"
)

type requestUsecase struct {
	requestRepositories repositories.Repositories
	validator           validator.Validator
}

func NewRequestUsecase(requestRepositories repositories.Repositories, validator validator.Validator) UseCase {
	return &requestUsecase{
		requestRepositories: requestRepositories,
		validator:           validator,
	}
}
