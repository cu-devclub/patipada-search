package handlers

import (
	"auth-service/users/usecases"
)

type usersHttpHandler struct {
	usersUsecase usecases.UsersUsecase
}

func NewUsersHttpHandler(usersUsecase usecases.UsersUsecase) UsersHandler {
	return &usersHttpHandler{
		usersUsecase: usersUsecase,
	}
}
