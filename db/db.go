package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"os"
)


var DB *sql.DB

func InitDB() {
	var err error
	
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	
	if port == "" {
		port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host+":"+port, database)

	
		

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Error opening database: ", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		age INT NOT NULL
		);`

	_, err = DB.Exec(createTable)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully")

}
