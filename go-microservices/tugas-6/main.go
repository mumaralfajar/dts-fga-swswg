package main

import "dts-fga-swswg/go-microservices/tugas-6/routers"

func main() {
	var PORT = ":8080"

	routers.StartServer().Run(PORT)
}