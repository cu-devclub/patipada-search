package repositories

import "auth-service/users/entities"

// InsertUsersData inserts users data into the database.
//
// It takes an `InsertUsersDto` object as a parameter and returns an error.
func (r *usersPostgresRepository) InsertUsersData(in *entities.InsertUsersDto) error {
	data := &entities.Users{
		Id:        in.Id,
		Username:  in.Username,
		Password:  in.Password,
		Salt:      in.Salt,
		Email:     in.Email,
		Role:      in.Role,
		Is_Active: true,
	}

	result := r.db.Create(data)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

