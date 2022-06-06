package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	APP_ENV    string
	PORT       string
	DEBUG      bool
	MONGO_HOST string
	MONGO_PORT string
	MOGO_DB    string
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
