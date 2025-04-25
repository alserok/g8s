package middleware

import "net/http"

type Middleware func(handler http.Handler) http.HandlerFunc

func With(handler http.Handler, mws ...Middleware) {
	for _, mw := range mws {
		handler = mw(handler)
	}
}
