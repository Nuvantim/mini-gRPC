package main

import (
	"example/config"      
	"example/database"    
	"example/service"     
	"log"

	"example/server"
)

func main() {
	// Inisialisasi database
	database.InitDB()
	defer database.CloseDB()

	// Buat instance layanan yang diperlukan
	categoryService := &service.CategoryService{}

	// Inisialisasi server dengan layanan yang dibutuhkan
	// Kami meneruskan categoryService ke fungsi NewServer
	port := os.Getenv("PORT")
	appServer := server.NewServer(":"+port, categoryService)

	// Mulai server di goroutine terpisah
	go func() {
		log.Println("Server starting on :"+port)
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Menangani graceful shutdown
	config.GracefulShutdown(appServer.Server) // Mengakses http.Server dari struct Server Anda
	log.Println("Server stopped")
}