package main

import (
	"log"
	"micro/database"
	"micro/config"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	pb "micro/generated"
)

type Server struct {
	pb.UnimplementedCategoryServiceServer
}

// Register the gRPC service
func RegisterServe(serv *grpc.Server) {
	pb.RegisterCategoryServiceServer(serv, Server{})
}

func main() {
	// Initialize PostgreSQL database connection
	database.InitDB()
	defer database.CloseDB()

	// Get service port from .env
	port := os.Getenv("PORT")

	// Create gRPC listener
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create gRPC server with logging interceptor
	serv := grpc.NewServer(
		grpc.UnaryInterceptor(config.LoggingInterceptor),
	)

	RegisterServe(serv)

	// Create a channel to catch shutdown signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	// Run gRPC server in a goroutine
	go func() {
		log.Println("Server is running on port " + port + "...")
		if err := serv.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Wait for shutdown signal
	<-stop
	log.Println("Shutting down server...")

	// Perform graceful shutdown
	serv.GracefulStop()

	// Wait for ongoing requests to complete
	time.Sleep(5 * time.Second)
	log.Println("Server stopped gracefully")
}
