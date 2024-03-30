package config

import (
	"github.com/spf13/viper"
)

var cfg *Config

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
		DataPath              string
		RecordPath            string
		LDAPath               string
		StopwordPath          string
		LogsPath              string
		SearchLogsDraftPath   string
		SearchLogsConfirmPath string
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
			DataPath:              viper.GetString("STATIC_DATA"),
			RecordPath:            viper.GetString("RECORD_DATA_PATH"),
			LDAPath:               viper.GetString("LDA_DATA_PATH"),
			StopwordPath:          viper.GetString("STOPWORD_PATH"),
			LogsPath:              viper.GetString("LOGS_PATH"),
			SearchLogsDraftPath:   viper.GetString("SEARCH_LOG_DRAFT_PATH"),
			SearchLogsConfirmPath: viper.GetString("SEARCH_LOG_CONFIRM_PATH"),
		},
	}
}

func GetConfig() Config {
	return *cfg
}
