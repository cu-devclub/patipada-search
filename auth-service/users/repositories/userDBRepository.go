package repositories

import "auth-service/users/entities"

type UsersRepository interface {
	InsertUsersData(in *entities.InsertUsersDto) error
	GetAllUsersData() ([]*entities.Users, error)
	GetUserByUsername(username string) (*entities.Users, error)
	UpdateUser(in *entities.Users) error
	RemoveUser(id string) error
}
