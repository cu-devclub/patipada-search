package main

import (
	"log/slog"
	"ml-gateway-service/config"
	"ml-gateway-service/logging"
	"ml-gateway-service/server"
)

func main() {
	logging.NewSLogger()

	config.ReadConfig()
	cfg := config.GetConfig()
	slog.Info("Reading Config successfully")

	s := server.NewGinServer(&cfg)

	go server.GRPCListen(&s, &cfg)

	s.Start()
}
