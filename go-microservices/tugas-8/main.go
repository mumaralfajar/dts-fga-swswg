package main

import (
	"dts-fga-swswg/go-microservices/tugas-8/database"
	"dts-fga-swswg/go-microservices/tugas-8/routers"
)

func main() {
	var PORT = ":8080"
	database.StartDB()
	routers.StartServer().Run(PORT)
}