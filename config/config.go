package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DB       DB
	Security Security
	HTTP     HTTPConfig
}

type DB struct {
	URI string `env:"DB_URI" envDefault:"postgresql://postgres:password@localhost:5555/auth"`
}

type Security struct {
	SecretKey string `env:"SECRET_KEY"`
}

type HTTPConfig struct {
	Port int `env:"HTTP_PORT" envDefault:"8089"`
}

func Parse() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("load config from env: %w", err)
	}
	return &cfg, nil
}
