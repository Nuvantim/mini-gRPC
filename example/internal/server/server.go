package server

import (
	"example/config"                     
	"example/rpc/proto/category/v1/categoryconnect"
	"example/internal/service"                     
	"net/http"

	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	*http.Server
}

// NewServer membuat dan mengembalikan instance Server yang sudah dikonfigurasi.
// Ini menerima alamat dan layanan yang akan di-serve.
func NewServer(addr string, categoryService *service.CategoryService) *Server {
	mux := http.NewServeMux()

	// Daftarkan handler gRPC-Connect Anda
	// Path dan handler diambil dari generated code
	path, handler := categoryconnect.NewCategoryServiceHandler(categoryService)
	mux.Handle(path, handler)

	// Konfigurasi CORS middleware
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Izinkan semua origin untuk development
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"*"}, // Izinkan semua header
		AllowCredentials: true,
	}).Handler(h2c.NewHandler(config.LogRequest(mux), &http2.Server{}))

	// Buat instance http.Server
	s := &http.Server{
		Addr:    addr,
		Handler: corsHandler, // Gunakan handler dengan CORS dan h2c
	}

	return &Server{s}
}