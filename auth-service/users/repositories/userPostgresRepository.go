package repositories

import (
	"auth-service/users/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type usersPostgresRepository struct {
	db *gorm.DB
}

func NewUsersPostgresRepository(db *gorm.DB) UsersRepository {
	return &usersPostgresRepository{db: db}
}

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
		log.Errorf("InsertUsersData: %v", result.Error)
		return result.Error
	}

	log.Debugf("InsertUsersData: %v", result.RowsAffected)
	return nil
}

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
