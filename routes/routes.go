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
	r.Use(middleware.RecoverMiddleware)
	r.Use(middleware.TimeoutMiddleware)
	r.Use(middleware.RateLimitMiddleware)
	// Test Cases
	r.HandleFunc("/panic", handlers.PanicTest).Methods("GET")
	r.HandleFunc("/slow", handlers.SlowHandler).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	
	// Protected Routes
	users := r.PathPrefix("/users").Subrouter()
	users.Use(middleware.JWTMiddleware)
	users.HandleFunc("", handlers.GetUsers).Methods("GET")
	users.HandleFunc("/{id}", handlers.GetUser).Methods("GET")
	users.HandleFunc("/{id}", handlers.UpdateUser).Methods("PUT")
	users.HandleFunc("", handlers.CreateUser).Methods("POST")
	users.HandleFunc("{id}", handlers.DeleteUser).Methods("DELETE")
	return r
}
