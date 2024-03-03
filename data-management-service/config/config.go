package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      App
		RabbitMQ RabbitMQ
		DB       Database
	}
	App struct {
		Port          int
		GRPCPort      int
		FrontendURL   string
		AuthService   string
		SearchService string
	}
	RabbitMQ struct {
		URL      string
		Username string
		Password string
	}
	Database struct {
		Host     string
		Port     string
		Username string
		Password string
		Dbname   string
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
		log.Printf("Error reading config file, %s", err)
	}
}

func GetConfig() Config {
	log.Println("Getting config...")
	return Config{
		App: App{
			Port:          viper.GetInt("APP_PORT"),
			GRPCPort:      viper.GetInt("GRPC_PORT"),
			FrontendURL:   viper.GetString("FRONTEND_URL"),
			AuthService:   viper.GetString("AUTH_SERVICE"),
			SearchService: viper.GetString("SEARCH_SERVICE"),
		},
		RabbitMQ: RabbitMQ{
			URL:      viper.GetString("RABBITMQ_URL"),
			Username: viper.GetString("RABBITMQ_USERNAME"),
			Password: viper.GetString("RABBITMQ_PASSWORD"),
		},
		DB: Database{
			Host:     viper.GetString("MONGO_DB_HOST"),
			Port:     viper.GetString("MONGO_DB_PORT"),
			Username: viper.GetString("MONGO_DB_USERNAME"),
			Password: viper.GetString("MONGO_DB_PASSWORD"),
			Dbname:   viper.GetString("MONGO_DB_NAME"),
		},
	}
}
