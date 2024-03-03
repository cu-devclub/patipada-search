package main

import (
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	"data-management/server"
	validator "data-management/structValidator"
	"log"
)
// TODO : make service can run even other service not there (gRPC connection)
func main() {
	/// ----------------- Initialize Config----------------- ///
	config.InitializeViper("./")
	cfg := config.GetConfig()
	log.Println("Getting Config success!")
	/// ----------------- Initialize Config----------------- ///

	/// ----------------- Initialize Database----------------- ///
	db := database.NewMongoDatabase(&cfg)
	// db := database.NewMockMongoDatabase()
	/// ----------------- Initialize Database----------------- ///

	validate := validator.NewValidator()

	/// ----------------- Initialize Communication----------------- ///

	grpc := communication.NewgRPC(&cfg)
	// grpc := communication.NewMockgRPC()

	rabbit, err := communication.ConnectToRabbitMQ(&cfg)
	if err != nil {
		log.Println("Error connecting to RabbitMQ", err)
		return
	}
	defer rabbit.Conn.Close()

	comm := communication.NewCommunicationImpl(*grpc, *rabbit)

	/// ----------------- Initialized Communication----------------- ///

	log.Println("Starting server...")
	server.NewGinServer(&cfg, &db, &validate, comm).Start()
}
