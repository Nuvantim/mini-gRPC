package main

import (
	"example/config"
	"example/database"
	"example/internal/server"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	port := os.Getenv("PORT_SERVICE")
	if port == "" {
		log.Fatal("PORT_SERVICE not set")
	}

	srv := server.New(":" + port)

	// Start server in goroutine
	go func() {
		log.Printf("Server is running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Setup graceful shutdown
	config.GracefulShutdown(srv.Server)
	defer database.CloseDB()
}
