package middleware 

import (
	"net/http"

)

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
