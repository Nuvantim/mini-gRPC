package server

import (
	"context"
	"example/config"
	"fmt"
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	*http.Server
}

type HandlerRegistrator func() (path string, handler http.Handler)

func New(addr string,corsOpts cors.Options,services ...HandlerRegistrator,) (*Server, error) {
	mux := http.NewServeMux()

	// Registrasi handler dengan validasi
	for _, register := range services {
		path, handler := register()
		if path == "" || handler == nil {
			return nil, fmt.Errorf("invalid handler registration")
		}
		mux.Handle(path, handler)
	}

	// Middleware chain: CORS → Logging → h2c
	handlerChain := cors.New(corsOpts).Handler(
		config.LogRequest(mux), // Pastikan LogRequest aman untuk concurrent access
	)

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: h2c.NewHandler(handlerChain, &http2.Server{}),
		},
	}, nil
}

// Shutdown untuk graceful termination
func (s *Server) Shutdown(ctx context.Context) error {
	return s.Server.Shutdown(ctx)
}
