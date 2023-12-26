package repositories

import "auth-service/users/entities"

type UsersRepository interface {
	InsertUsersData(in *entities.InsertUsersDto) error
	GetAllUsersData() ([]*entities.Users, error)
	GetUserByUsername(username string) (*entities.Users, error)
	UpdateUser(in *entities.Users) error
	// ResetPassword(in *entities.ResetPassword) error
	// StoreResetToken(userID string, in *entities.ResetToken) error
	RemoveUser(username string) error
}
