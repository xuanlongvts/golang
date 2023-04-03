package config

import (
	"errors"
	"github.com/spf13/viper"
)

var EnvBuild string // important, this var has to declare global to get environment variable

func init() {
	viper.SetConfigFile("./config/dev.json")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic(errors.New("config file not found"))
		}
		panic(errors.New("can not load config file"))
	}
}

func LoadConfig() string {
	return EnvBuild
}
