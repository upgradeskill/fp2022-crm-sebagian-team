package config

import (
	"time"

	"github.com/spf13/viper"
)

// JWTConfig represents JWT configuration.
type JWTConfig struct {
	AccessSecret  string
	RefreshSecret string
	TTL           time.Duration
}

// LoadJWTConfig loads JWT configuration.
func LoadJWTConfig() JWTConfig {
	return JWTConfig{
		AccessSecret: viper.GetString("JWT_ACCESS_SECRET"),
		TTL:          viper.GetDuration("JWT_TTL"),
	}
}
