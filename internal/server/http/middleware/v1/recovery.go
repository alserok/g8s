package v1

import (
	"net/http"

	"github.com/alserok/g8s/internal/utils/logger"
)

func WithRecovery(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.ExtractContext(r.Context()).Error("recovered from panic", logger.WithArg("error", err))
			}
		}()

		handler.ServeHTTP(w, r)
	}
}
