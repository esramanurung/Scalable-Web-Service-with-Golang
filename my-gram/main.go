package main

import (
	"my-gram/databases"
	"my-gram/routers"
)

func main() {
	databases.StartDB()
	r := routers.StartApp()

	r.Run(":8080")
}
