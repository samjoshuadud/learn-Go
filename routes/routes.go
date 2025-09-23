package routes

import (
	"github.com/gorilla/mux"
	"learnGo/handlers"
	"learnGo/middleware"

)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	
	// Middlewares
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.AuthMiddleware)
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.TimeoutMiddleware)
	r.Use(middleware.RateLimitMiddleware)
	// Test Cases
	r.HandleFunc("/panic", handlers.PanicTest).Methods("GET")
	r.HandleFunc("/slow", handlers.SlowHandler).Methods("GET")

	// Endpoints
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", handlers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	return r
}
