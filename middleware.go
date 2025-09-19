package main

import (
	"fmt"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("-> %s %s\n", r.Method, r.RequestURI)

		next.ServeHTTP(w, r)

		fmt.Printf("Finished handling %s %s\n", r.Method, r.RequestURI)
	})

	
}
