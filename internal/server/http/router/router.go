package router

import (
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
)

func SetupRoutes(mux *http.ServeMux, handler handler.Handler) {
	handle(mux, "GET /pods/:namespace", handler.GetPods)
}

func handle(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Handle(path, v1.WithErrorHandler(handler))
	}
}
