package main

import (
	"fmt"
	database "jokenpo/database"
	"os"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	app := &api{
		config: config{
			addr: os.Getenv("STATISTICS_ADDR"),
			port: ":" + os.Getenv("STATISTICS_PORT"),
		},
	}
	if err := app.run(); err != nil {
		fmt.Fprintf(os.Stderr, "Server failed to start: %v\n", err)
		os.Exit(1)
	}
}
