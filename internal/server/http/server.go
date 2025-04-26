package http

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserok/g8s/internal/server/http/handler"
	"github.com/alserok/g8s/internal/server/http/middleware"
	v1 "github.com/alserok/g8s/internal/server/http/middleware/v1"
	"github.com/alserok/g8s/internal/server/http/router"
	"github.com/alserok/g8s/internal/service"
	"github.com/alserok/g8s/internal/utils/logger"
	"net/http"
	"time"
)

func New(srvc service.Service, log logger.Logger) *server {
	mux := http.NewServeMux()
	srvr := &http.Server{
		Handler:      mux,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
	}

	router.SetupRoutes(mux, handler.New(srvc))

	middleware.With(srvr.Handler, v1.WithLogger(log), v1.WithRecovery)

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
	s.srvr.Addr = fmt.Sprintf(":%s", port)

	if err := s.srvr.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to serve: " + err.Error())
	}
}
