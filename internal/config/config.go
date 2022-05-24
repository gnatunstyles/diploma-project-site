package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
	DBConnString string `env:"DB_CONFIG_STRING"`
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config initialization error: %s", err)
	}
	// log.Info().Msgf("config: %v", cfg)
	return cfg, nil
}
