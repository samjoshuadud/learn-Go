package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)


var DB *sql.DB

func InitDB() {
	var err error
	dsn := "go_user:password@tcp(127.0.0.1:3306)/go_users"

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
