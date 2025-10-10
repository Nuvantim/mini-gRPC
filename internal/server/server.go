package server

import (
	"log"
	"net/http"
	"time"

	"example/config"
	"example/database"
	"example/internal/repository"
	"example/internal/service"
	"example/middleware"
	"example/rpc/proto/category/v1/categoryconnect"
	"example/rpc/proto/product/v1/productconnect"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"golang.org/x/time/rate"
)

type Server struct {
	*http.Server
}

func New() *Server {
	database.InitDB()
	queries := repository.New(database.DB)

	// Define ConnectRPC
	categoryService := service.NewCategoryService(queries)
	productService := service.NewProductService(queries)

	mux := http.NewServeMux()
	mux.Handle(categoryconnect.NewCategoryServiceHandler(categoryService))
	mux.Handle(productconnect.NewProductServiceHandler(productService))

	// Load Server Configuration
	var serv, err = config.GetServerConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Limiter Configuration
	rateLimiterConfig := middleware.RateLimiterConfig{
		Rate:         rate.Every(time.Duration(serv.Rate) * time.Minute), // Waiting time for the next request
		Burst:        serv.Burst,                                         // The number of request allowed
		PerClient:    true,                                               // Limit per client IP
		LRUCacheSize: serv.LRU,                                           // LRU cache capacity limit
	}

	// Build middleware chain
	middlewareChain := middleware.Chain(
		middleware.CORS(),
		middleware.RateLimiter(rateLimiterConfig),
		middleware.Logging(),
	)

	// Integration http2
	handler := middlewareChain(mux)
	h2cHandler := h2c.NewHandler(handler, &http2.Server{})

	// Define port service
	return &Server{
		Server: &http.Server{
			Addr:    ":" + serv.Port,
			Handler: h2cHandler,
                        ReadHeaderTimeout: 5 * time.Second,
		},
	}
}
