package server

import (
	"net/http"

	"example/config"
	"example/database"
	"example/internal/repository"
	"example/internal/service"
	"example/rpc/proto/category/v1/categoryconnect"
	//"example/pb/proto/product/v1/productv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	*http.Server
}

func New(addr string) *Server {
	// Initialize dependencies
	database.InitDB()
	queries := repository.New(database.DB)

	// Create services
	categoryService := service.NewCategoryService(queries)
	//productService := product.NewProductService(queries)

	// Setup handlers with logging middleware
	mux := http.NewServeMux()
	mux.Handle(categoryconnect.NewCategoryServiceHandler(categoryService))
	// mux.Handle(productconnect.NewProductServiceHandler(productService))

	// Apply middleware chain
	handler := config.LogRequest(mux)

	// Configure HTTP/2
	h2cHandler := h2c.NewHandler(handler, &http2.Server{})

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: h2cHandler,
		},
	}
}
