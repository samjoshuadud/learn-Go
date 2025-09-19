package main 

import (
	"fmt"
	"log"
	"net/http"
	"learnGo/db"
	"learnGo/routes"
)

func main() {
	db.InitDB()

	r := routes.SetupRoutes()

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))	
}
