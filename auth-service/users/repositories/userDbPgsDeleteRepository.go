package repositories

import "auth-service/users/entities"

func (r *usersPostgresRepository) RemoveUser(username string) error {
	result := r.db.Where("username = ?", username).Delete(&entities.Users{})
	return result.Error
}
