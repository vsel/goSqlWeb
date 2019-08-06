package service

import (
	"log"

	"github.com/spf13/viper"

	configStruct "github.com/vsel/goSqlWeb/config/struct"
)

// GetConfig from yml
func GetConfig(prefix string) configStruct.Configuration {
	viper.SetConfigName("config")
	viper.AddConfigPath(prefix + "/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read the config file: %s", err)
	}

	var configuration configStruct.Configuration

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Failed to parse configuration file: %v", err)
	}

	return configuration
}
