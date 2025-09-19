package middleware 

import (
	"fmt"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("-> %s %s\n", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)

		fmt.Printf("Finished handling %s %s\n", r.Method, r.RequestURI)
	})
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

			if apiKey != "123" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return 
			}

		next.ServeHTTP(w, r)
	})
}
