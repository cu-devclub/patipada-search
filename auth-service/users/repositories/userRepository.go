package repositories

import "auth-service/users/entities"

type UsersRepository interface {
	InsertUsersData(in *entities.InsertUsersDto) error
	GetAllUsersData() ([]*entities.Users, error)
}
