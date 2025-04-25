package v1

import (
	"net/http"
)

func WithRecovery(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != "" {
				// TODO
			}
		}()

		handler.ServeHTTP(w, r)
	}
}
