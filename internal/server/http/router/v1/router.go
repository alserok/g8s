package v1

import (
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
)

func SetupRoutes(mux *http.ServeMux, h handler.Handler) {
	handle(mux, "GET /deployments/{namespace}", h.V1.GetDeployments)
	handle(mux, "POST /deployments/", h.V1.CreateDeployment)
}

func handle(mux *http.ServeMux, path string, handler func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mux.Handle(path, v1.WithErrorHandler(handler))
	}
}
