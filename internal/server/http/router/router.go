package router

import (
	"github.com/alserok/g8s/internal/cache"
	"net/http"

	_ "github.com/alserok/g8s/docs"
	"github.com/alserok/g8s/internal/server/http/handler"
	v1 "github.com/alserok/g8s/internal/server/http/router/v1"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SetupRoutes(mux *http.ServeMux, handler handler.Handler, cache cache.Repository) {
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	v1.SetupRoutes(mux, handler, cache)
}
