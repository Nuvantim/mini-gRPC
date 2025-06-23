package main

import (
	"example/config"
	"example/database"
	"example/internal/server"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Chechk env
	godotenv.Load()
	port := os.Getenv("PORT_SERVICE")
	if port == "" {
		log.Fatalf("Port Service Not Found")
	}

	srv := server.New(":" + port)

	// Start server in goroutine for graceful shutdown
	go func() {
		log.Printf("Server running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Setup graceful shutdown
	config.GracefulShutdown(srv.Server)
	defer database.CloseDB()
}
