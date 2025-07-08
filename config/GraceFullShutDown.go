package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// GracefulShutdown handles server shutdown gracefully
func GracefulShutdown(srv *http.Server) {
	// Siapkan channel untuk menerima signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Menunggu signal shutdown
	sigReceived := <-quit
	log.Printf("Shutdown signal received: %v", sigReceived)

	// Gunakan context dengan timeout untuk memberi waktu server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Pastikan cancel selalu dipanggil untuk membersihkan resource context

	// Matikan server dengan graceful shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Server exited gracefully")

	// Setelah server shutdown, kita bisa menutup channel quit
	close(quit)
}
