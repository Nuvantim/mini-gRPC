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
	// Chechk env
	godotenv.Load()
	port := os.Getenv("PORT_SERVICE")
	if port == "" {
		log.Fatalf("Port Service Not Found")
	}

	srv := server.New(":" + port)

	// Make Channel for Error Server
	serverErr := make(chan error, 1)

	// Start server in goroutine for graceful shutdown
	go func() {
		log.Printf("Server running on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			serverErr <- err
		}
	}()

	// Setup graceful shutdown
	config.GracefulShutdown(srv, serverErr)
	defer database.CloseDB()
}
