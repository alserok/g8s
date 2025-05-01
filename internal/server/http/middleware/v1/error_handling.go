package v1

import (
	"github.com/alserok/g8s/internal/utils/errors"
	"github.com/alserok/g8s/internal/utils/logger"
	"net/http"
)

func WithErrorHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			loc, msg, code := errors.Parse(errors.HTTP, err)
			logger.ExtractContext(r.Context()).Error(err.Error(), logger.WithArg("loc", loc))
			w.WriteHeader(code)
			w.Write([]byte(msg))
		}
	}
}
