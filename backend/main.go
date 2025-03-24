package backend

import (
	"fmt"
	"jokenpo/database"
)

func main() {
	database.InitDB()
	//Close it when it is not used
	defer database.CloseDB()

	db := database.GetDB()
	fmt.Println("Using database... ", db)
}
