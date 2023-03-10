package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Prefix       int8 `envconfig:"PREFIX"`
	Limit        int  `envconfig:"LIMIT"`
	TimeCooldown int  `envconfig:"TIME_COOLDOWN"`
}

func New() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("[config.New]:error loading %v: %v", ".env", err)
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("[config.New]:can't process envs: %v", err)
	}
	return &cfg, nil

}
