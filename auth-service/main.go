package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/server"
	usersMigrate "auth-service/users/migrations"
	"fmt"
	"log"
)

func main() {
	// TODO : Implementing log service 
	log.Println("Initializing config")
	config.InitializeViper("./")
	cfg := config.GetConfig()

	db := database.NewPostgresDatabase(&cfg)
	log.Println("Success connect to database")

	err := usersMigrate.UsersMigrate(db)
	if err != nil {
		_ = fmt.Errorf("failed to migrate %w", err)
		return
	}

	s := server.NewEchoServer(&cfg, db.GetDb())

	go server.GRPCListen(s,&cfg)

	s.Start()

}
