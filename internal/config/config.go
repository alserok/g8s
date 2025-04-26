package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Env  string
	Port string

	KubeConfigPath string `json:"kube_config"`
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.KubeConfigPath = os.Getenv("KUBECONFIG_PATH")

	return &cfg
}
