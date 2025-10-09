package main

import (
	"log"
	"net/http"

	"example/config"
	"example/database"
	"example/internal/server"
)

func main() {
	// Check Environtment
	conf, err := config.CheckEnv()
	if err != nil {
		log.Fatal(err)
	}

	// Start Server
	srv := server.New()

	// print Banner
	config.Banner()
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
