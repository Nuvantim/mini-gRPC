package middleware

import (
	"log"
	"net/http"
	"path"
	"time"
)

type Middleware func(http.Handler) http.Handler

func Logging() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				clientIP := r.RemoteAddr
				userAgent := r.UserAgent()
				endpoint := path.Base(r.RequestURI)

				log.Printf(
					"%s %s %s %v | Client IP: %s | User-Agent: %s",
					r.Method,
					endpoint,
					time.Since(start),
					clientIP,
					userAgent,
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
