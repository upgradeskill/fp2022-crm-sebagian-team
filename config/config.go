package config

import (
	"log"

	"github.com/spf13/viper"
)

// Config is the struct for the config
type Config struct {
	JWT JWTConfig
}

// NewConfig creates a new Config struct
func NewConfig() *Config {
	viper.SetConfigFile(`.env`)
	viper.AutomaticEnv()
	viper.ReadInConfig()

	if viper.GetBool(`DEBUG`) {
		log.Println("Service RUN on DEBUG mode")
	}

	return &Config{
		JWT: LoadJWTConfig(),
	}
}
