package server

import (
	"net/http"

	"example/middleware"
	"example/database"
	"example/internal/repository"
	"example/internal/service"
	"example/rpc/proto/category/v1/categoryconnect"
	"example/rpc/proto/product/v1/productconnect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Server struct {
	*http.Server
}

func New(addr string) *Server {
	database.InitDB()
	queries := repository.New(database.DB)

	categoryService := service.NewCategoryService(queries)
	productService := service.NewProductService(queries)

	mux := http.NewServeMux()
	mux.Handle(categoryconnect.NewCategoryServiceHandler(categoryService))
	mux.Handle(productconnect.NewProductServiceHandler(productService))


	// Build middleware chain
	middlewareChain := middleware.Chain(
		middleware.CORS(),
		// middleware.CSRF(),
		middleware.Logging(),
	)

	handler := middlewareChain(mux)
	h2cHandler := h2c.NewHandler(handler, &http2.Server{})

	return &Server{
		Server: &http.Server{
			Addr:    addr,
			Handler: h2cHandler,
		},
	}
}