package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

type Config struct {
	Env  string
	Port string

	AI struct {
		ApiToken string `json:"api_token"`
	}

	Cache struct {
		Addr string        `json:"host"`
		TTL  time.Duration `json:"ttl"`
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.AI.ApiToken = os.Getenv("AI_API_TOKEN")

	cfg.Cache.Addr = os.Getenv("CACHE_ADDR")
	ttl, err := time.ParseDuration(os.Getenv("CACHE_TTL"))
	if err == nil {
		cfg.Cache.TTL = ttl
	} else {
		cfg.Cache.TTL = time.Hour * 24
	}

	return &cfg
}
