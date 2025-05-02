package v1

import (
	"github.com/alserok/g8s/internal/cache"
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
)

func SetupRoutes(mux *http.ServeMux, h handler.Handler, cache cache.Repository) {
	handle(mux, "GET /v1/list/{namespace}", v1.WithCache(h.V1.List, func(r *http.Request) string { return r.PathValue("namespace") }, cache))
	handle(mux, "POST /v1/create", h.V1.Create)
	handle(mux, "DELETE /v1/delete", h.V1.Delete)
	handle(mux, "PUT /v1/update", h.V1.Update)
}

func handle(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) {
	mux.HandleFunc(path, v1.WithErrorHandler(handler))
}
