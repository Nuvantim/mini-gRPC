package config
import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// Unary interceptor for logging incoming requests and errors
func LoggingInterceptor(ctx context.Context,req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler,) (interface{}, error) {
	// Get client address if available
	p, _ := peer.FromContext(ctx)
	clientAddr := "unknown"
	if p != nil {
		clientAddr = p.Addr.String()
	}

	// Log the incoming request
	log.Printf("Incoming request: %s from %s", info.FullMethod, clientAddr)

	// Handle the request
	resp, err := handler(ctx, req)

	// Log error if occurs
	if err != nil {
		log.Printf("Error handling request %s from %s: %v", info.FullMethod, clientAddr, err)
	}

	return resp, err
}