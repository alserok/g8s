package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Env  string
	Port string
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	return &cfg
}
