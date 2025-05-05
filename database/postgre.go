package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

/*
	The functions are captalized making them PUBLIC (and passible to be exported) thats important

since we are with 2 different go projects (on on the backend an another one on the database and in
the future another one on the websocket)
*/

func InitDB() {
	serviceURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
	)
	fmt.Println("DB_HOST from env:", os.Getenv("DB_HOST"))

	fmt.Println("Waiting for the database Startup")

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	var err error
	maxRetries := 5
	backoff := time.Second

	for i := 1; i <= maxRetries; i++ {
		fmt.Printf("Attempt %d: Connecting to database...\n", i)
		db, err = sql.Open("postgres", conn.String())
		if err != nil {
			log.Printf("Error opening DB: %v", err)
		} else {
			err = db.Ping()
			if err == nil {
				fmt.Println("Database connection established.")
				break
			}
			log.Printf("Ping failed: %v", err)
		}

		if i == maxRetries {
			log.Fatalf("Could not connect to database after %d attempts", maxRetries)
		}

		waitTime := backoff * time.Duration(i)
		fmt.Printf("Waiting %s before retrying...\n", waitTime)
		time.Sleep(waitTime)
	}

	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var result string
		err = rows.Scan(&result)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Version: %s\n", result)
	}
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
