//! This file is used for testing purposes only. It is used to isolate the service from the actual gRPC and RabbitMQ services.
//! The actual service is implemented in the communication package.
//! By mocking the gRPC and RabbitMQ services, we can test the service without having to rely on the actual services.

package main

import (
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	"data-management/logging"
	mock "data-management/mock/communication"
	"data-management/request/migration"
	"data-management/server"
	validator "data-management/structValidator"
	"log/slog"
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

	db, err := database.NewMongoDatabase(&cfg)
	if err != nil {
		slog.Error("Failed to connect to database", slog.String("err", err.Error()))
		return
	}
	slog.Info("Connect to db successfully!")

	validate := validator.NewValidator()

	grpc := mock.NewMockgRPC()
	slog.Info("Connect to gRPC successfully!")

	rabbit := mock.MockRabbitMQ()
	slog.Info("Connect to RabbitMQ successfully!")

	comm := communication.NewCommunicationImpl(grpc, rabbit)

	serv := server.NewGinServer(&cfg, &db, &validate, &comm)
	slog.Info("Server initialized successfully!")

	if err = migration.Migration(&cfg, &db, &serv); err != nil {
		slog.Error("failed to migrate %w", err)
		return
	}
	slog.Info("Migration successfully!")

	serv.Start()
}
