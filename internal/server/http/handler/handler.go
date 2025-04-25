package handler

import (
	v1 "github.com/alserok/g8s/internal/server/http/handler/v1"
	"github.com/alserok/g8s/internal/service"
)

func New(service service.Service) Handler {
	return Handler{
		v1.Handler{Service: service},
		service,
	}
}

type Handler struct {
	v1.Handler

	service service.Service
}
