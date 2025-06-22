package main

import (
	"example/database"
	"example/internal/server"
	"example/internal/service"
	"log"
)

func main() {
	db := database.InitDB()
	defer db.Close()
	port := os.Getenv("PORT")
	srv := server.New(
		":"+port,
		category.New(db.Queries).Register, // Langsung pass method
		product.New(db.Queries).Register,  // Tanpa interface
	)

	log.Println("Server starting on :"+port)
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	config.GracefulShutdown(srv.Server)
	log.Println("Server stopped gracefully")
}
