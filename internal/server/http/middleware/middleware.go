package middleware

import "net/http"

type Middleware func(handler http.Handler) http.HandlerFunc

func With(handler http.Handler, mws ...Middleware) http.Handler {
	for _, mw := range mws {
		handler = mw(handler)
	}

	return handler
}
