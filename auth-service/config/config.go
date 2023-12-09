package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App App
		Db  Db
	}

	App struct {
		Port       int
		JWTKey     string
		SuperAdmin SuperAdmin
		RolesMap   map[string]int
	}

	SuperAdmin struct {
		Username string
		Password string
		Role     string
		Email    string
	}

	Db struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
		TimeZone string
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
	
	// Load roles map from config
	rolesMap := make(map[string]int)
	if err := viper.UnmarshalKey("app.rolesMap", &rolesMap); err != nil {
		panic(fmt.Errorf("error loading roles map: %w", err))
	}

	return Config{
		App: App{
			Port:   viper.GetInt("app.server.port"),
			JWTKey: viper.GetString("app.jwt.key"),
			SuperAdmin: SuperAdmin{
				Username: viper.GetString("app.super-admin.username"),
				Password: viper.GetString("app.super-admin.password"),
				Role:     viper.GetString("app.super-admin.role"),
				Email:    viper.GetString("app.super-admin.email"),
			},
			RolesMap: rolesMap,
		},
		Db: Db{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.user"),
			Password: viper.GetString("database.password"),
			DBName:   viper.GetString("database.dbname"),
			SSLMode:  viper.GetString("database.sslmode"),
			TimeZone: viper.GetString("database.timezone"),
		},
	}
}

