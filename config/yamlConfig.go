package config

import (
	"github.com/spf13/viper"
)

func YamlConfig() {
	viper.AddConfigPath("config")
	viper.SetConfigName("application")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Can't connect to yaml")
	}
}
