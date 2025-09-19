package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")	
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	
	r.Use(loggingMiddleware)

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))	
}
