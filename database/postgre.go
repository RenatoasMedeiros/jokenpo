package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func InitDB() {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println("Waiting for the database Startup")

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		err = db.Ping()
		if err == nil {
			fmt.Println("Successfully connected to the database")
			return
		}
	}
	fmt.Println("Successfully connected to the database")
}

func GetDB() *sql.DB {
	if db == nil {
		log.Fatalf("Database connection has not been initialized. Call InitDB first.")
	}
	return db
}

func CloseDB() {
	if db != nil {
		if err := db.Close(); err != nil {
			log.Fatalf("Error closing the database: %v", err)
		} else {
			fmt.Println("Database connection closed successfully.")
		}
	}
}
