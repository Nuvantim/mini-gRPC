package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"connectrpc.com/connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "example/gen/greet/v1"
	"example/gen/greet/v1/greetv1connect"
)

type (
	GreetReq = connect.Request[greetv1.GreetRequest]
	GreetRes = connect.Response[greetv1.GreetResponse]
)

// LoggingMiddleware membungkus handler dengan logging
func LoggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Incoming %s %s", r.Method, r.URL.Path)
		
		h.ServeHTTP(w, r)
		
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}

type GreetServer struct{}

func (s *GreetServer) Greet(ctx context.Context, req *GreetReq) (*GreetRes, error) {
	log.Printf("Processing Greet request for: %s", req.Msg.Name)
	
	res := &GreetRes{
		Msg: &greetv1.GreetResponse{
			Greeting: "Hello, " + req.Msg.Name + "!",
		},
	}
	
	log.Printf("Greet response generated for: %s", req.Msg.Name)
	return res, nil
}

func main() {
	// Setup logger dengan prefix dan flags
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("Initializing Connect RPC server...")

	// Buat router dengan middleware logging
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(&GreetServer{})
	mux.Handle(path, LoggingMiddleware(handler))

	// Konfigurasi server
	server := &http.Server{
		Addr:    ":8080",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	// Log server startup
	log.Printf("Server starting on %s", server.Addr)
	
	// Start server dengan error handling
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
