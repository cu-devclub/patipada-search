package main

import (
	"auth-service/config"
	"auth-service/database"
	"auth-service/logging"
	"auth-service/server"
	usersMigrate "auth-service/users/migrations"
	"log/slog"
)

func main() {
	logging.NewSLogger()

	if err := config.InitializeViper("./"); err != nil {
		slog.Error("failed to initialize viper %w", err)
		return
	}
	slog.Info("Viper initialized successfully")

	config.ReadConfig()
	cfg := config.GetConfig()

	db, err := database.NewPostgresDatabase(&cfg)
	if err != nil {
		slog.Error("failed to connect to database %w", err)
		return
	}
	slog.Info("Successfully connected to database!")

	if err := usersMigrate.UsersMigrate(db); err != nil {
		slog.Error("failed to migrate default user %w", err)
		return
	}

	s := server.NewEchoServer(&cfg, db.GetDb())

	go server.GRPCListen(s, &cfg)

	s.Start()
}
