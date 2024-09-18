package main

import (
	"update-data/database"
	"update-data/router"
)

func main() {
	database.StartDB()

	r := router.StartApp()

	r.Run(":8080")
}
