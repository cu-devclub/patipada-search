package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App   App
		Db    Db
		Email Email
		Link  Link
		User  User
	}

	App struct {
		Port        int
		GRPCPort	int
		FrontendURL string
		JWTKey      string
		RolesMap    map[string]int
	}

	User struct {
		SuperAdmin UserCredential
		Admins     UserCredential
		Users      UserCredential
	}

	UserCredential struct {
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
	Email struct {
		Host              string
		Port              int
		SenderName        string
		SenderEmail       string
		SenderPassword    string
		ReceiverTestEmail string
	}

	Link struct {
		URL string
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
	// Load roles map from config
	rolesMap := make(map[string]int)
	superAdmin := viper.GetInt("ROLES_MAP_SUPER_ADMIN")
	admin := viper.GetInt("ROLES_MAP_ADMIN")
	user := viper.GetInt("ROLES_MAP_USER")
	rolesMap["super-admin"] = superAdmin
	rolesMap["admin"] = admin
	rolesMap["user"] = user

	return Config{
		App: App{
			Port:        viper.GetInt("SERVER_PORT"),
			GRPCPort:    viper.GetInt("GRPC_PORT"),
			FrontendURL: viper.GetString("FRONTEND_URL"),
			JWTKey:      viper.GetString("JWT_KEY"),
			RolesMap:    rolesMap,
		},
		User: User{
			SuperAdmin: UserCredential{
				Username: viper.GetString("SUPER_ADMIN_USERNAME"),
				Password: viper.GetString("SUPER_ADMIN_PASSWORD"),
				Role:     viper.GetString("SUPER_ADMIN_ROLE"),
				Email:    viper.GetString("SUPER_ADMIN_EMAIL"),
			},
			Admins: UserCredential{
				Username: viper.GetString("ADMIN_USERNAME"),
				Password: viper.GetString("ADMIN_PASSWORD"),
				Role:     viper.GetString("ADMIN_ROLE"),
				Email:    viper.GetString("ADMIN_EMAIL"),
			},
			Users: UserCredential{
				Username: viper.GetString("USER_USERNAME"),
				Password: viper.GetString("USER_PASSWORD"),
				Role:     viper.GetString("USER_ROLE"),
				Email:    viper.GetString("USER_EMAIL"),
			},
		},
		Db: Db{
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetInt("DATABASE_PORT"),
			User:     viper.GetString("DATABASE_USER"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			DBName:   viper.GetString("DATABASE_DBNAME"),
			SSLMode:  viper.GetString("DATABASE_SSLMODE"),
			TimeZone: viper.GetString("DATABASE_TIMEZONE"),
		},
		Email: Email{
			Host:              viper.GetString("EMAIL_SERVER_HOST"),
			Port:              viper.GetInt("EMAIL_SERVER_PORT"),
			SenderName:        viper.GetString("EMAIL_SENDER_NAME"),
			SenderEmail:       viper.GetString("EMAIL_SENDER_EMAIL"),
			SenderPassword:    viper.GetString("EMAIL_SENDER_PASSWORD"),
			ReceiverTestEmail: viper.GetString("EMAIL_RECEIVER_TESTEMAIL"),
		},
		Link: Link{
			URL: viper.GetString("LINK_URL"),
		},
	}
}
