package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Redis struct {
		DNS      string
		Password string
	}
	MySQL struct {
		DSN string
	}
}

func Get(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return nil, err
	}

	cfg := Config{}

	cfg.Redis.DNS = os.Getenv("REDIS_DSN")
	cfg.MySQL.DSN = os.Getenv("MYSQL_DSN")

	return &cfg, nil
}
