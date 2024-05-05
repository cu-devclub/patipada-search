//! This file is used to isolate the main function for testing purposes
//! The actual service is implemented in the communication package
//! By mocking the gRPC service, we can test the service without having to rely on the actual service

package main

import (
	"log/slog"
	"search-esdb-service/communication"
	"search-esdb-service/config"
	"search-esdb-service/data"
	"search-esdb-service/database"
	"search-esdb-service/logging"
	mock "search-esdb-service/mock/communication"
	"search-esdb-service/monitoring"
	recordMigrator "search-esdb-service/record/migration"
	"search-esdb-service/server"
)

func main() {

	logging.NewSLogger()

	if err := config.InitializeViper("./"); err != nil {
		slog.Error("failed to initialize viper %w", err)
		return
	}
	slog.Info("Viper initialized successfully!")

	config.ReadConfig()
	cfg := config.GetConfig()

	cfg.ReadMlConfig()

	// If the usecase is bigger, this one can be an object
	// right now it used only to set up the prometheus counter
	monitoring.NewMonitoring()

	if err := data.ReadStopWord(&cfg); err != nil {
		slog.Error("Failed to read stop word", slog.String("err", err.Error()))
		return
	}
	slog.Info("Read stop word successfully!")

	db, err := database.NewElasticDatabase(&cfg)
	if err != nil {
		slog.Error("Failed to connect to database", slog.String("err", err.Error()))
		return
	}
	slog.Info("Connect to es db successfully!")

	// db,err := database.MockElasticDatabase(&cfg)
	// if err != nil {
	// 	slog.Error("Failed to connect to database", slog.String("err", err.Error()))
	// 	return
	// }

	err = recordMigrator.MigrateRecords(&cfg, db)
	if err != nil {
		slog.Error("Failed to migrate records", slog.String("err", err.Error()))
		return
	}
	slog.Info("Migrate records successfully!")

	grpc := mock.NewMockgRPC()
	slog.Info("Connect to gRPC successfully!")
	comm := communication.NewCommunicationImpl(grpc)

	s := server.NewGinServer(&cfg, db.GetDB(), &comm)

	s.Start()

}
