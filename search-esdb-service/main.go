package main

import (
	"log/slog"
	"search-esdb-service/communication"
	"search-esdb-service/config"
	"search-esdb-service/constant"
	"search-esdb-service/data"
	"search-esdb-service/database"
	"search-esdb-service/logging"
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

	err = recordMigrator.MigrateRecords(&cfg, db)
	if err != nil {
		slog.Error("Failed to migrate records", slog.String("err", err.Error()))
		return
	}
	slog.Info("Migrate records successfully!")

	s := server.NewGinServer(&cfg, db.GetDB())

	rabbitMQ, err := communication.ConnectToRabbitMQ(&cfg, db.GetDB(), *s.GetRecordArch())
	if err != nil {
		slog.Error("Failed to connect to RabbitMQ", slog.String("err", err.Error()))
		return
	}
	slog.Info("Connect to RabbitMQ successfully!")

	comm := communication.NewCommunicationImpl(*rabbitMQ)

	go func() {
		err := comm.Listen([]string{constant.UPDATE_RECORD_TOPIC})
		if err != nil {
			slog.Error("Failed to listen to RabbitMQ", slog.String("err", err.Error()))
		}
	}()
	go server.GRPCListen(s, &cfg)

	s.Start()

}
