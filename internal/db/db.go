package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
)

// DB is the database connection
var DB *sql.DB

// InitDB initializes the mysql database connection
func InitDB() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Set the maximum number of open connections to the database
	DB.SetMaxOpenConns(10)

	// Set the maximum number of idle connections to the database
	DB.SetMaxIdleConns(5)

	// Set the maximum lifetime of a connection to the database
	DB.SetConnMaxLifetime(time.Minute * 5)

	// Ping the database to ensure the connection is working
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to database")
}
