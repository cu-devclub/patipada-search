package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/server"
	usersMigrate "auth-service/users/migrations"
	"fmt"
)

func main() {
	config.InitializeViper("./")
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)

	err := usersMigrate.UsersMigrate(db)
	if err != nil {
		_ = fmt.Errorf("failed to migrate %w", err)
		return 
	}

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}
