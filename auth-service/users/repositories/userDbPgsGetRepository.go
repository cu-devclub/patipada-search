package repositories

import "auth-service/users/entities"

// GetAllUsersData retrieves all users data from the repository.
//
// It returns a slice of pointers to entities.Users and an error, if any.
func (r *usersPostgresRepository) GetAllUsersData() ([]*entities.Users, error) {
	users := make([]*entities.Users, 0)
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *usersPostgresRepository) GetUserByUsername(username string) (*entities.Users, error) {
	user := &entities.Users{}
	if err := r.db.Where("username = ?", username).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}
