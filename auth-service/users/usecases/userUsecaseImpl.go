package usecases

import (
	"auth-service/users/repositories"
)

type UsersUsecaseImpl struct {
	usersRepository repositories.UsersRepository
	UserEmailing    repositories.UserEmailing
}

func NewUsersUsecaseImpl(usersRepository repositories.UsersRepository, userEmailing repositories.UserEmailing) UsersUsecase {
	return &UsersUsecaseImpl{
		usersRepository: usersRepository,
		UserEmailing:    userEmailing,
	}
}
