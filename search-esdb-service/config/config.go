package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App
		ESDB     ESDB
		RabbitMQ RabbitMQ
		Static   Static
	}

	App struct {
		Port        int
		FrontendURL string
		GRPCPort    int
	}
	ESDB struct {
		URL      string
		Username string
		Password string
	}
	RabbitMQ struct {
		URL      string
		Username string
		Password string
	}
	Static struct {
		DataPath     string
		RecordPath   string
		LDAPath      string
		StopwordPath string
	}
)

func InitializeViper(path string) {
	log.Println("Initializing viper...")
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
	log.Println("Viper initialized")
}

func GetConfig() Config {
	log.Println("Getting config...")
	return Config{
		App: App{
			Port:        viper.GetInt("SERVER_PORT"),
			FrontendURL: viper.GetString("FRONTEND_URL"),
			GRPCPort:    viper.GetInt("GRPC_PORT"),
		},
		ESDB: ESDB{
			URL:      viper.GetString("ESDB_URL"),
			Username: viper.GetString("ESDB_USERNAME"),
			Password: viper.GetString("ESDB_PASSWORD"),
		},
		RabbitMQ: RabbitMQ{
			URL:      viper.GetString("RABBITMQ_URL"),
			Username: viper.GetString("RABBITMQ_USERNAME"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		Static: Static{
			DataPath:     viper.GetString("STATIC_DATA"),
			RecordPath:   viper.GetString("RECORD_DATA_PATH"),
			LDAPath:      viper.GetString("LDA_DATA_PATH"),
			StopwordPath: viper.GetString("STOPWORD_PATH"),
		},
	}
}
