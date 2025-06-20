package main

import (
	"example/config"      // Pastikan path ini benar
	"example/database"    // Pastikan path ini benar
	"example/service"     // Pastikan path ini benar
	"log"

	// Import server package yang baru
	"your_project/server" // Ganti 'your_project' dengan nama modul Anda
)

func main() {
	// Inisialisasi database
	if _, err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Buat instance layanan yang diperlukan
	categoryService := &service.CategoryService{}

	// Inisialisasi server dengan layanan yang dibutuhkan
	// Kami meneruskan categoryService ke fungsi NewServer
	appServer := server.NewServer(":8080", categoryService)

	// Mulai server di goroutine terpisah
	go func() {
		log.Println("Server starting on :8080")
		if err := appServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Menangani graceful shutdown
	config.GracefulShutdown(appServer.Server) // Mengakses http.Server dari struct Server Anda
	log.Println("Server stopped")
}