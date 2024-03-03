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
	log.Println("Initializing config...")
	config.InitializeViper("./")
	log.Println("Getting config...")
	cfg := config.GetConfig()
	log.Println("Config initialized!")
	/// ----------------- Initialize Config----------------- ///

	/// ----------------- Initialize Database----------------- ///
	log.Println("Connecting to MongoDB...")
	db := database.NewMongoDatabase(&cfg)
	// db := database.NewMockMongoDatabase()
	log.Println("Connected to MongoDB!")
	/// ----------------- Initialize Database----------------- ///

	validate := validator.NewValidator()

	/// ----------------- Initialize Communication----------------- ///
	log.Println("Connecting to gRPC...")
	grpc := communication.NewgRPC(&cfg)
	// grpc := communication.NewMockgRPC()
	log.Println("Connected to gRPC!")

	log.Println("Connecting to RabbitMQ...")
	rabbit, err := communication.ConnectToRabbitMQ(&cfg)
	if err != nil {
		log.Println("Error connecting to RabbitMQ", err)
		return
	}
	defer rabbit.Conn.Close()
	log.Println("Connected to RabbitMQ!")

	comm := communication.NewCommunicationImpl(*grpc, *rabbit)
	// err = testPublishRabbitMQ(comm)
	// if err != nil {
	// 	log.Println("Error testing RabbitMQ", err)
	// 	return
	// }
	/// ----------------- Initialized Communication----------------- ///

	log.Println("Starting server...")
	server.NewGinServer(&cfg, &db, &validate, comm).Start()
}

// func testPublishRabbitMQ(comm communication.Communication) error {
// 	log.Println("Publishing to RabbitMQ...")
// 	entity := &entities.UpdateRecord{
// 		DocumentID: "60d5ecf7c88f9a200f9e2c5a",
// 		Question:   "Updated question",
// 		Answer:     "Updated answer",
// 		StartTime:  "2021-06-25T00:00:00Z",
// 		EndTime:    "2021-06-25T00:00:10Z",
// 	}

// 	err := comm.PublishUpdateRecordsToRabbitMQ(constant.UPDATE_RECORD_PAYLOAD_NAME, entity)
// 	if err != nil {
// 		log.Println("Error publishing to RabbitMQ", err)
// 		return err
// 	}
// 	log.Println("Published to RabbitMQ!!!!")
// 	return nil
// }
