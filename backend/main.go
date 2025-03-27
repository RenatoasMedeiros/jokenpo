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
			addr: ":" + os.Getenv("BACKEND_PORT"),
			//addr: ":8080",
		},
	}
	app.run()
}
