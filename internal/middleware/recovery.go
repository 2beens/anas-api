package middleware

import (
	"log"
	"net/http"
	"runtime/debug"
)

func PanicRecovery() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(respWriter http.ResponseWriter, req *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("http: panic serving %s: %v\n%s", req.URL.Path, r, debug.Stack())
				}
			}()

			// handler call
			next.ServeHTTP(respWriter, req)
		})
	}
}
