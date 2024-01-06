package main

import (
	"data-management/config"
	"data-management/database"
	"data-management/server"
	validator "data-management/structValidator"
	"log"
)

func main() {
	// TODO : Implementing token authorization
	log.Println("Initializing config...")
	config.InitializeViper("./")
	log.Println("Getting config...")
	cfg := config.GetConfig()

	log.Println("Connecting to MongoDB...")
	db := database.NewMongoDatabase(&cfg)
	log.Println("Connected to MongoDB!")

	validate := validator.NewValidator()

	log.Println("Starting server...")
	server.NewGinServer(&cfg, &db, &validate).Start()
}
