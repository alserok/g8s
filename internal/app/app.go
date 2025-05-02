package app

import (
	"github.com/alserok/g8s/internal/cache/memcached"
	"os"
	"os/signal"
	"syscall"

	"github.com/alserok/g8s/internal/config"
	"github.com/alserok/g8s/internal/external/ai"
	"github.com/alserok/g8s/internal/external/k8s"
	"github.com/alserok/g8s/internal/metrics"
	"github.com/alserok/g8s/internal/server"
	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/utils/logger"
)

func MustServe(cfg *config.Config) {
	log := logger.New(logger.Slog, cfg.Env)
	defer func() {
		_ = log.Close()
	}()

	log.Info("starting app 👾")

	log.Info("initializing clients ☄️")

	k8sClient := k8s.NewClient()
	log.Info("k8s client initialized 🛸")
	aiClient := ai.NewClient(cfg.AI.ApiToken)
	log.Info("ai client initialized 👽")

	log.Info("initializing metrics 🌌")
	metr := metrics.New()
	log.Info("metrics initialized 🪐")

	log.Info("initializing layers 🚀")

	srvc := service.New(k8sClient, aiClient, metr)
	log.Info("service initialized 🧬")

	log.Info("initializing cache 👾")
	c := memcached.New(cfg.Cache.Addr, cfg.Cache.TTL)
	defer func() {
		_ = c.Close()
	}()
	log.Info("cache initialized 🙂")

	srvr := server.New(server.HTTP, srvc, metr, c, log)
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
