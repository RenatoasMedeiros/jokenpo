package main

import (
	database "jokenpo/database"
	"os"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	app := &api{
		config: config{
			addr: os.Getenv("BACKEND_ADDR"),
			port: ":" + os.Getenv("BACKEND_PORT"),
		},
	}
	app.run()
}
