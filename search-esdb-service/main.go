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

	/// ----------------- Initialize config ----------------- ///
	config.InitializeViper("./")
	cfg := config.GetConfig()
	log.Println("Config initialized!")
	/// ----------------- Initialized config ----------------- ///

	/// ----------------- Initialize database ----------------- ///
	db := database.NewElasticDatabase(&cfg)
	/// ----------------- Initialized database ----------------- ///

	/// ----------------- Migrate record ----------------- ///
	recordMigrator.MigrateRecords(&cfg, db)
	/// ----------------- Migrated record ----------------- ///

	/// ----------------- Initialize communication ----------------- ///
	rabbitMQ, err := communication.ConnectToRabbitMQ(&cfg, db.GetDB())
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err)
		return
	}

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
	go server.GRPCListen(s, &cfg)

	log.Println("Starting HTTP server on Port", cfg.App.Port, "...")
	s.Start()
	/// ----------------- Started server ----------------- ///
}
