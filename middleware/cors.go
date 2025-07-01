package middleware

import (
    "net/http"
    
    "github.com/rs/cors"
)

func CORS() Middleware {
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{"GET", "POST"},
        AllowedHeaders:   []string{"Content-Type", "Connect-Protocol-Version","Connect-Timeout","Grpc-Timeout","X-Grpc-Web","X-User-Agent"},
        // AllowedHeaders:   []string{"Accept", "Content-Type", "X-CSRF-Token"},
        ExposedHeaders:   []string{"Grpc-Status","Grpc-Message", "Grpc-Status-Details-Bin"},
        AllowCredentials: true,
        MaxAge:           7200,
    })
    
    return func(next http.Handler) http.Handler {
        return c.Handler(next)
    }
}