package middleware

import (
	"log"
	"net"
	"net/http"

	lru "github.com/hashicorp/golang-lru/v2"
	"golang.org/x/time/rate"
)

// RateLimiterConfig holds configuration for the rate limiter.
type RateLimiterConfig struct {
	Rate rate.Limit
	Burst int
	PerClient bool
	LRUCacheSize int
}

// getClientIP extracts the client's IP address from the request.
func getClientIP(r *http.Request) string {
	if ip, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
		return ip
	}
	return "unknown"
}

// RateLimiter creates a new HTTP middleware for rate limiting.
func RateLimiter(config RateLimiterConfig) func(next http.Handler) http.Handler {
	var globalLimiter *rate.Limiter
	var clientLimitersCache *lru.Cache[string, *rate.Limiter] // LRU cache untuk limiter per klien

	if !config.PerClient {
		globalLimiter = rate.NewLimiter(config.Rate, config.Burst)
	} else {
		if config.LRUCacheSize <= 0 {
			// Atur ukuran default jika tidak valid
			config.LRUCacheSize = 1000 
		}
		var err error
		// Buat cache LRU baru. Anda bisa menambahkan OnEvict untuk melihat apa yang dihapus.
		clientLimitersCache, err = lru.New[string, *rate.Limiter](config.LRUCacheSize)
		if err != nil {
			// Ini adalah error fatal saat inisialisasi middleware
			panic(log.printf("ERROR: Gagal membuat LRU cache: %v", err))
		}
		log.Printf("LRU Cache Size: %d\n", config.LRUCacheSize)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var limiter *rate.Limiter

			if config.PerClient {
				ip := getClientIP(r)
				
				// Coba dapatkan limiter dari cache
				if val, ok := clientLimitersCache.Get(ip); ok {
					limiter = val
				} else {
					// Jika tidak ada di cache, buat yang baru dan tambahkan ke cache
					newLimiter := rate.NewLimiter(config.Rate, config.Burst)
					clientLimitersCache.Add(ip, newLimiter)
					limiter = newLimiter
				}
			} else {
				limiter = globalLimiter
			}

			if limiter.Allow() {
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
				return
			}
		})
	}
}
