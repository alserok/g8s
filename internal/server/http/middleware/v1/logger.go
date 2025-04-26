package v1

import (
	"github.com/alserok/g8s/internal/utils/logger"
	"net/http"
)

func WithLogger(log logger.Logger) func(handlerFunc http.Handler) http.HandlerFunc {
	return func(fn http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(logger.WrapContext(r.Context(), log))

			fn.ServeHTTP(w, r)
		}
	}
}
