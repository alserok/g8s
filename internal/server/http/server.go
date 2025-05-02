package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	_ "github.com/swaggo/http-swagger"

	"github.com/alserok/g8s/docs"
	"github.com/alserok/g8s/internal/server/http/handler"
	"github.com/alserok/g8s/internal/server/http/middleware"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
	"github.com/alserok/g8s/internal/server/http/router"
	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/utils/logger"
)

func New(srvc service.Service, log logger.Logger) *server {
	mux := http.NewServeMux()

	router.SetupRoutes(mux, handler.New(srvc))

	srvr := &http.Server{
		Handler:      middleware.With(mux, v1.WithHeaders, v1.WithRecovery, v1.WithLogger(log)),
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 15,
	}

	return &server{
		srvr: srvr,
	}
}

type server struct {
	srvr *http.Server
}

func (s server) Shutdown() error {
	return s.srvr.Shutdown(context.Background())
}

func (s server) MustServe(port string) {
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%s", port)

	s.srvr.Addr = fmt.Sprintf(":%s", port)

	if err := s.srvr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to serve: " + err.Error())
	}
}
