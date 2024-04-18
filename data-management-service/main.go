package main

import (
	"data-management/communication"
	"data-management/config"
	"data-management/database"
	"data-management/logging"
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

	grpc, err := communication.NewgRPC(&cfg)
	if err != nil {
		slog.Error("Failed to connect to gRPC", slog.String("err", err.Error()))
		return
	}
	// grpc := communication.NewMockgRPC()
	slog.Info("Connect to gRPC successfully!")

	rabbit, err := communication.ConnectToRabbitMQ(&cfg)
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", slog.String("err", err.Error()))
		return
	}
	defer rabbit.Conn.Close()
	// rabbit := communication.MockRabbitMQ()
	slog.Info("Connect to RabbitMQ successfully!")

	comm := communication.NewCommunicationImpl(*grpc, *rabbit)

	serv := server.NewGinServer(&cfg, &db, &validate, &comm)
	slog.Info("Server initialized successfully!")

	if err = migration.Migration(&cfg, &db, &serv); err != nil {
		slog.Error("failed to migrate %w", err)
		return
	}
	slog.Info("Migration successfully!")

	serv.Start()
}