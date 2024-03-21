package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

type (
	Config struct {
		App      App
		RabbitMQ RabbitMQ
		DB       Database
	}
	App struct {
		Port           int
		GRPCPort       int
		FrontendURL    string
		AuthService    string
		AuthGRPCPort   int
		SearchService  string
		SearchGRPCPort int
		DataSourcePath string
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

func InitializeViper(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func ReadConfig() {
	cfg = &Config{
		App: App{
			Port:           viper.GetInt("APP_PORT"),
			GRPCPort:       viper.GetInt("GRPC_PORT"),
			FrontendURL:    viper.GetString("FRONTEND_URL"),
			AuthService:    viper.GetString("AUTH_SERVICE"),
			AuthGRPCPort:   viper.GetInt("AUTH_GRPC_PORT"),
			SearchService:  viper.GetString("SEARCH_SERVICE"),
			SearchGRPCPort: viper.GetInt("SEARCH_GRPC_PORT"),
			DataSourcePath: viper.GetString("STATIC_DATA"),
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

func GetConfig() Config {
	return *cfg
}
