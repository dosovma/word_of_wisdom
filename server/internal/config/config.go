package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host    string `envconfig:"SERVER_HOST" default:"127.0.0.1"`
	Port    string `envconfig:"SERVER_PORT" default:":9000"`
	Timeout int    `envconfig:"SERVER_TIMEOUT" default:"5"`
}

func NewServer() (*Config, error) {
	cfg := &Config{}

	err := envconfig.Process("SERVER", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch envs: %w", err)
	}

	return cfg, nil
}
