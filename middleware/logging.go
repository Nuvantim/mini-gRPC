package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Printf(
					"%s %s %s %v",
					r.Method,
					r.RequestURI,
					r.RemoteAddr,
					time.Since(start),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
