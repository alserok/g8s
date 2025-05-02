package v1

import (
	"net/http"
)

func WithHeaders(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn.ServeHTTP(w, r)
		w.Header().Add("Content-Type", "application/json")
	}
}
