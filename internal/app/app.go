package app

import (
	"github.com/alserok/g8s/internal/config"
	"github.com/alserok/g8s/internal/external/ai"
	"github.com/alserok/g8s/internal/external/k8s"
	"github.com/alserok/g8s/internal/metrics"
	"github.com/alserok/g8s/internal/server"
	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/utils/logger"
	"os"
	"os/signal"
	"syscall"
)

func MustServe(cfg *config.Config) {
	log := logger.New(logger.Slog, cfg.Env)
	defer func() {
		_ = log.Close()
	}()

	log.Info("starting app 👾")

	log.Info("initializing clients ☄️")

	k8sClient := k8s.NewClient(cfg.K8s.KubeConfigPath)
	log.Info("k8s client initialized 🛸", logger.WithArg("kubeconfig_path", cfg.K8s.KubeConfigPath))
	aiClient := ai.NewClient(cfg.AI.ApiToken)
	log.Info("ai client initialized 👽")

	log.Info("initializing metrics 🌌")
	metr := metrics.New()
	log.Info("metrics initialized 🪐")

	log.Info("initializing layers 🚀")

	srvc := service.New(k8sClient, aiClient, metr)
	log.Info("service initialized 🧬")

	srvr := server.New(server.HTTP, srvc, metr, log)
	log.Info("server initialized 🧪")

	log.Info("app is serving ⚗️", logger.WithArg("port", cfg.Port))
	run(srvr, cfg.Port)
	log.Info("app was closed 🧫")
}

func run(server server.Server, port string) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

	go server.MustServe(port)

	<-ch

	_ = server.Shutdown()
}
