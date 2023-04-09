package main

import (
	"dts-fga-swswg/go-microservices/tugas-7/config"
	"dts-fga-swswg/go-microservices/tugas-7/routers"
)

func main() {
	config.ConnectDatabase()
	routers.StartServer().Run(":8080")
}