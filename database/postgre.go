package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

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

	fmt.Println("Waiting for the database Startup")

	conn, _ := url.Parse(serviceURI)
	conn.RawQuery = "sslmode=verify-ca;sslrootcert=ca.pem"

	var err error
	db, err = sql.Open("postgres", conn.String())

	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	rows, err := db.Query("SELECT version()")
	if err != nil {
		panic(err)
	}

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
