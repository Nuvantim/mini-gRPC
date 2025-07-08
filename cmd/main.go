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
	env, err := cofig.CheckEnv()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(env)

	serv := config.ServerEnvirontment()

	srv := server.New(":"+serv.port, serv.rate, serv.burst, serv.lru)

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
