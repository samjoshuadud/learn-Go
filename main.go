package main 

import (
	"fmt"
	"log"
	"net/http"
	"learnGo/db"
	"learnGo/routes"

	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file")
	}
	db.InitDB()

	r := routes.SetupRoutes()
	port := os.Getenv("PORT")	

	fmt.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))	
}
