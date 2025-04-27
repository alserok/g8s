package v1

import (
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
)

func SetupRoutes(mux *http.ServeMux, h handler.Handler) {
	handle(mux, "GET /list/{namespace}", h.V1.List)
	handle(mux, "POST /create/", h.V1.Create)
	handle(mux, "POST /delete/", h.V1.Delete)
	handle(mux, "POST /update/", h.V1.Update)
}

func handle(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Handle(path, v1.WithErrorHandler(handler))
	}
}
