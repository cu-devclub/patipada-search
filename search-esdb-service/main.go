package main

import (
	"log"
	"search-esdb-service/config"
	"search-esdb-service/database"
	recordMigrator "search-esdb-service/record/migration"
	"search-esdb-service/server"
)

func main() {
	log.Println("Starting server...")

	log.Println("Initializing config...")
	config.InitializeViper("./")

	cfg := config.GetConfig()
	log.Println("Config initialized:", cfg)

	log.Println("Connecting to database...")
	db := database.NewElasticDatabase(&cfg)
	log.Println("Success connect to database:")

	log.Println("Starting migration...")
	recordMigrator.RecordMigrate(&cfg, db)
	log.Println("Migration finished")

	s := server.NewGinServer(&cfg, db.GetDB())

	log.Println("Starting gRPC server...")
	go server.GRPCListen(s, &cfg)

	log.Println("Starting HTTP server on Port", cfg.App.Port, "...")
	s.Start()

}
