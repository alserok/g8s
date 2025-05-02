package server

import (
	"github.com/alserok/g8s/internal/cache"
	"github.com/alserok/g8s/internal/metrics"
	"github.com/alserok/g8s/internal/server/http"
	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/utils/logger"
)

const (
	HTTP = iota
)

type Server interface {
	MustServe(port string)
	Shutdown() error
}

func New(t uint, srvc service.Service, metr metrics.Metrics, cache cache.Repository, log logger.Logger) Server {
	switch t {
	case HTTP:
		return http.New(srvc, cache, log)
	default:
		panic("invalid server type")
	}
}
