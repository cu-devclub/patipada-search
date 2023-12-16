package migrations

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/users/entities"
	"auth-service/users/helper"

	"gorm.io/gorm"
)

// UsersMigrate migrates the users in the database.
//
// It takes a database instance as the parameter.
// It returns an error if there was an issue during the migration process.
func UsersMigrate(db database.Database) error {
	// 1. Get super admin username and password from env
	cfg := config.GetConfig()
	// 2. Check if super admin already exists
	users, err := getAllUsers(db.GetDb())
	if err != nil {
		return err
	}
	user := helper.GetUserFromUserLists(users, cfg.App.SuperAdmin.Username)
	if user != nil {
		//* alrady have super admin in database 
		return nil  
	}

	// 3. Insert super admin 
	uuid,err := helper.GenerateUUID()
	if err != nil {
		return err
	}
	password,salt,err := helper.GenerateHashedSaltedPassword(cfg.App.SuperAdmin.Password)
	if err != nil {
		return err
	}

	superAdmin := &entities.Users{
		Id:       uuid,
		Username: cfg.App.SuperAdmin.Username,
		Password: password,
		Salt:     salt,
		Email:    cfg.App.SuperAdmin.Email,
		Role:     cfg.App.SuperAdmin.Role,
		Is_Active: true,
	}

	if err := insertUser(db.GetDb(), superAdmin); err != nil {
		return err
	}

	return nil 
}

// getAllUsers retrieves all users from the database.
//
// It takes a *gorm.DB as a parameter and returns a slice of *entities.Users and an error.
func getAllUsers(db *gorm.DB) ([]*entities.Users, error) {
	users := make([]*entities.Users, 0)
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// insertUser inserts a user into the database.
//
// Parameters:
// - db: A pointer to a gorm.DB object representing the database connection.
// - user: A pointer to a entities.Users object representing the user to be inserted.
//
// Returns:
// - error: An error object representing any error that occurred during the insertion process.
func insertUser(db *gorm.DB, user *entities.Users) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}