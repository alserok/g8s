package handler

import (
	v1 "github.com/alserok/g8s/internal/server/http/handler/v1"
	"github.com/alserok/g8s/internal/service"
)

func New(service service.Service) Handler {
	return Handler{
		V1: v1.Handler{Service: service},
	}
}

type Handler struct {
	V1 v1.Handler
}
