package middleware

import (
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"net/http"
)

func CORS() Middleware {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"example.com"},
		AllowedMethods:   connectcors.AllowedMethods(),
		AllowedHeaders:   connectcors.AllowedHeaders(),
		ExposedHeaders:   connectcors.ExposedHeaders(),
		AllowCredentials: true,
		MaxAge:           7200,
	})

	return func(next http.Handler) http.Handler {
		return c.Handler(next)
	}
}
