package main

import (
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	"data-management/server"
	validator "data-management/structValidator"
	"log"
)

func main() {
	log.Println("Initializing config...")
	config.InitializeViper("./")
	log.Println("Getting config...")
	cfg := config.GetConfig()

	log.Println("Connecting to MongoDB...")
	db := database.NewMongoDatabase(&cfg)
	log.Println("Connected to MongoDB!")

	validate := validator.NewValidator()

	comm := communication.NewgRPC(&cfg)
	// comm := communication.NewTempgRPC()

	log.Println("Starting server...")
	server.NewGinServer(&cfg, &db, &validate, comm).Start()
}
