package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	Initialize()
}

// Initialize the configuration
func Initialize() {
	viper.SetDefault("port", 80)

	viper.AutomaticEnv()

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("../../..")

	if err := viper.ReadInConfig(); err != nil {
		logrus.Info(err)
	}

}
