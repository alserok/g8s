package config

import (
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type Config struct {
	Env  string
	Port string

	K8s struct {
		KubeConfigPath string `json:"kube_config"`
	}

	AI struct {
		ApiToken string `json:"api_token"`
	}
}

func MustLoad() *Config {
	var cfg Config

	cfg.Env = os.Getenv("ENV")
	cfg.Port = os.Getenv("PORT")

	cfg.K8s.KubeConfigPath = os.Getenv("KUBECONFIG_PATH")

	cfg.AI.ApiToken = os.Getenv("AI_API_TOKEN")

	return &cfg
}
