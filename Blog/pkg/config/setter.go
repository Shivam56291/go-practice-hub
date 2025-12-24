package config

import (
	"log"

	"github.com/spf13/viper"
)

func Set() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("Config file not found")
		} else {
			log.Fatal("Error reading config file:", err)
		}
	}

	if err := viper.Unmarshal(&configurations); err != nil {
		log.Fatal("Error unmarshalling config:", err)
	}
}