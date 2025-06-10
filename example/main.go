package main

import (
	"connectrpc.com/connect"
	"context"
	"example/config"
	"example/database"
	"example/gen/category/v1/categoryv1connect"
	"example/service"
	"log"
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	// Initialize database
	if _, err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Create services
	categoryService := &service.CategoryService{}

	// Setup HTTP server
	mux := http.NewServeMux()
	path, handler := categoryv1connect.NewCategoryServiceHandler(categoryService)
	mux.Handle(path, handler)

	// Add CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(h2c.NewHandler(config.LogRequest(mux), &http2.Server{}))

	// Start server
	server := &http.Server{
		Addr:    ":8080",
		Handler: corsHandler,
	}

	go func() {
		log.Println("Server starting on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Graceful shutdown
	config.GracefulShutdown(server)
	log.Println("Server stopped")
}
