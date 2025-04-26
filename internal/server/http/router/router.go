package router

import (
	"net/http"

	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/router/v1"
)

func SetupRoutes(mux *http.ServeMux, handler handler.Handler) {
	v1.SetupRoutes(mux, handler)
}
