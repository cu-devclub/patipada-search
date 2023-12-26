package repositories

import (
	"gorm.io/gorm"
)

type usersPostgresRepository struct {
	db *gorm.DB
}

func NewUsersPostgresRepository(db *gorm.DB) UsersRepository {
	return &usersPostgresRepository{db: db}
}
