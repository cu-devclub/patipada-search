package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App    App
		ESDB   ESDB
		Static Static
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
	Static struct {
		DataPath     string
		RecordPath   string
		StopWordPath string
	}
)

func InitializeViper(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
}

func GetConfig() Config {
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
		Static: Static{
			DataPath:     viper.GetString("STATIC_DATA"),
			RecordPath:   viper.GetString("RECORD_DATA_PATH"),
			StopWordPath: viper.GetString("STOPWORD_PATH"),
		},
	}
}
