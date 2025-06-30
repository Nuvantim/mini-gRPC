package middleware

import (
    "net/http"
    
    "github.com/rs/cors"
)

func CORS() Middleware {
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Link"},
        AllowCredentials: true,
        MaxAge:           300,
    })
    
    return func(next http.Handler) http.Handler {
        return c.Handler(next)
    }
}