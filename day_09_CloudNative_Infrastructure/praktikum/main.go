package main

import (
	"fmt"
	"log"

	"github.com/devianwahyu/efishery/api/router"
	"github.com/devianwahyu/efishery/database"
)

func main() {
	config := database.ConfigDB()

	db, err := database.NewDBConnection(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Database connected")

	e := router.InitRoute(db)
	e.Logger.Fatal(e.Start(":8080"))
}
