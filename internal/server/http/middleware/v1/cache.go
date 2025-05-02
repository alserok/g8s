package v1

import (
	"github.com/alserok/g8s/internal/cache"
	"net/http"
)

func WithCache(fn func(w http.ResponseWriter, r *http.Request) error, getKey func(r *http.Request) string, cache cache.Repository) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		b, err := cache.GetBytes(r.Context(), getKey(r))
		if err == nil {
			_, _ = w.Write(b)
			return nil
		}

		return fn(w, r)
	}
}
