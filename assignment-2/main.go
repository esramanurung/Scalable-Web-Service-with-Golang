package main

import (
	"assignment-2/databases"
	"assignment-2/routers"
)

func main() {
	databases.StartDB()
	routers.StartingServer().Run(":8080")
}
