package main

import (
	"log"
	"search-esdb-service/communication"
	"search-esdb-service/config"
	"search-esdb-service/constant"
	"search-esdb-service/database"
	recordMigrator "search-esdb-service/record/migration"
	"search-esdb-service/server"
)

func main() {
	log.Println("Starting server...")

	/// ----------------- Initialize config ----------------- ///
	log.Println("Initializing config...")
	config.InitializeViper("./")
	cfg := config.GetConfig()
	log.Println("Config initialized")
	/// ----------------- Initialized config ----------------- ///

	/// ----------------- Initialize database ----------------- ///
	log.Println("Connecting to database...")
	db := database.NewElasticDatabase(&cfg)
	log.Println("Success connect to database:")
	/// ----------------- Initialized database ----------------- ///

	/// ----------------- Migrate record ----------------- ///
	log.Println("Starting migration...")
	recordMigrator.MigrateRecords(&cfg, db)
	log.Println("Migration finished")
	/// ----------------- Migrated record ----------------- ///

	/// ----------------- Initialize communication ----------------- ///
	log.Println("Connecting to RabbitMQ...")
	rabbitMQ, err := communication.ConnectToRabbitMQ(&cfg, db.GetDB())
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return
	}
	log.Println("Success connect to RabbitMQ")

	comm := communication.NewCommunicationImpl(*rabbitMQ)
	go func() {
		err := comm.Listen([]string{constant.UPDATE_RECORD_TOPIC})
		if err != nil {
			log.Println("Failed to start listener:", err)
		}
	}()
	/// ----------------- Initialized communication ----------------- ///

	/// ----------------- Start server ----------------- ///
	s := server.NewGinServer(&cfg, db.GetDB())
	log.Println("Starting gRPC server...")
	go server.GRPCListen(s, &cfg)

	log.Println("Starting HTTP server on Port", cfg.App.Port, "...")
	s.Start()
	/// ----------------- Started server ----------------- ///
}
