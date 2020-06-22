package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT       string `envconfig:"app_api_port"`
	REDIS_PORT string `envconfig:"redis_port"`
}

func FromEnv() (*Config, error) {
	cfg := new(Config)

	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
