package repositories

import "auth-service/users/entities"


func (r *usersPostgresRepository) UpdateUser(in *entities.Users) error {
	result := r.db.Model(&entities.Users{}).Where("username = ?", in.Username).Updates(in)
	return result.Error
}