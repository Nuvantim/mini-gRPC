package middleware

import "net/http"

func CSRF() Middleware {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            // Skip for safe methods and Connect protocol
            if r.Method == http.MethodGet || 
               r.Method == http.MethodHead || 
               r.Method == http.MethodOptions ||
               r.Header.Get("Content-Type") == "application/connect+proto" {
                next.ServeHTTP(w, r)
                return
            }

            token := r.Header.Get("X-CSRF-Token")
            if token == "" {
                http.Error(w, "CSRF token missing", http.StatusForbidden)
                return
            }
            
            // Add proper token validation logic here
            next.ServeHTTP(w, r)
        })
    }
}