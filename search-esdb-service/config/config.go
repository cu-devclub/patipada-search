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
	}
	ESDB struct {
		URL      string
		Username string
		Password string
	}
	Static struct {
		DataPath string
	}
)

func GetConfig() Config {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %v", err))
	}
	//TODO : change docker environment name to match
	return Config{
		App: App{
			Port:        viper.GetInt("app.server.port"),
			FrontendURL: viper.GetString("app.frontend.url"),
		},
		ESDB: ESDB{
			URL:      viper.GetString("esdb.url"),
			Username: viper.GetString("esdb.username"),
			Password: viper.GetString("esdb.password"),
		},
		Static: Static{
			DataPath: viper.GetString("static.data"),
		},
	}
}
