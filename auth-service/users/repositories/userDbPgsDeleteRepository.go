package repositories

import "auth-service/users/entities"

func (r *usersPostgresRepository) RemoveUser(id string) error {
	result := r.db.Where("id = ?", id).Delete(&entities.Users{})
	return result.Error
}
