package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/server"
	usersMigrate "auth-service/users/migrations"
	"log"
)

func main() {
	config.InitializeViper("./")
	cfg := config.GetConfig()
	log.Println("Getting Config successfully....")

	db := database.NewPostgresDatabase(&cfg)

	err := usersMigrate.UsersMigrate(db)
	if err != nil {
		log.Println("failed to migrate %w", err)
		return
	}

	s := server.NewEchoServer(&cfg, db.GetDb())

	go server.GRPCListen(s, &cfg)

	log.Println("Starting HTTP server....")
	s.Start()
}
