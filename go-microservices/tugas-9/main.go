package main

import "time"

func main() {
	ticker := time.NewTicker(15 * time.Second) // make new ticker
	// call post request every 15second
	for ; true; <-ticker.C {
		PostRequest()
	}
}