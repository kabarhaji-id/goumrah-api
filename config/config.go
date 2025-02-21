package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS,required,notEmpty"`
	PostgresDSN   string `env:"POSTGRES_DSN,required,notEmpty"`
}

func New() (Config, error) {
	if err := godotenv.Load(); err != nil {
		return Config{}, err
	}

	return env.ParseAs[Config]()
}
