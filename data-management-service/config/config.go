package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		DB Database
	}
	App struct {
		Port int 
		FrontendURL string
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
	return Config{
		App: App{
			Port: viper.GetInt("APP_PORT"),
			FrontendURL: viper.GetString("FRONTEND_URL"),
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
