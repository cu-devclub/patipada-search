package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/server"
	usersMigrate "auth-service/users/migrations"
)

func main() {
	//TODO : Dockerized service 
	cfg := config.GetConfig()
	db := database.NewPostgresDatabase(&cfg)

	usersMigrate.UsersMigrate(db)

	server.NewEchoServer(&cfg, db.GetDb()).Start()
}

