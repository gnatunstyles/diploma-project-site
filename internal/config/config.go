package config

import (
	"crypto/tls"
	"diploma-project-site/internal/models"
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	JwtSecretKey string `env:"JWT_SECRET_KEY"`
	DBConnString string `env:"DB_CONFIG_STRING"`
	Tls          *tls.Config
}

func New() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("config initialization error: %s", err)
	}
	cer, err := tls.LoadX509KeyPair(models.CertPath, models.KeyPath)
	if err != nil {
		return nil, fmt.Errorf("config tls initialization error: %s", err)
	}
	cfg.Tls = &tls.Config{Certificates: []tls.Certificate{cer}}
	return cfg, nil
}
