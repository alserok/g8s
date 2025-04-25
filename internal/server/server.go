package server

import (
	"github.com/alserok/g8s/internal/server/http"
	"github.com/alserok/g8s/internal/service"
)

const (
	HTTP = iota
)

type Server interface {
	MustServe(port string)
	Shutdown() error
}

func New(t uint, srvc service.Service) Server {
	switch t {
	case HTTP:
		return http.New(srvc)
	default:
		panic("invalid server type")
	}
}
