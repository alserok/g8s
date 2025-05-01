package v1

import (
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
)

func SetupRoutes(mux *http.ServeMux, h handler.Handler) {
	handle(mux, "GET /v1/list/{namespace}", h.V1.List)
	handle(mux, "POST /v1/create", h.V1.Create)
	handle(mux, "POST /v1/delete", h.V1.Delete)
	handle(mux, "POST /v1/update", h.V1.Update)
}

func handle(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) {
	mux.HandleFunc(path, v1.WithErrorHandler(handler))
}
