package main

import (
	"os"
)

func main() {
	app := &api{
		config: config{
			addr: ":" + os.Getenv("BACKEND_PORT"),
		},
	}
	app.run()
}
