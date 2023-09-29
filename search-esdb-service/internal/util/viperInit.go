package util

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

func InitViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// TODO: To be decided?
			panic(fmt.Errorf("config file not found"))
		} else {
			panic(fmt.Errorf("error when reading config file: %s", err))
		}
	}
}
