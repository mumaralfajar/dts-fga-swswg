package main

import (
	"fmt"
	"sync"
)

func main() {
	s1 := []string{"bisa1", "bisa2", "bisa3"}
	s2 := []string{"coba1", "coba2", "coba3"}

	wg := sync.WaitGroup{}

	mutex := sync.Mutex{}

	for i := 1; i <= 8; i++ {
		wg.Add(1)
		mutex.Lock()
		if i%2 != 0 {
			go receive(i, s1, &wg, &mutex)
		} else {
			go receive(i, s2, &wg, &mutex)
		}
	}
	wg.Wait()
}

func receive(index int, data []string, wg *sync.WaitGroup, mutex *sync.Mutex) {
	fmt.Println(data, index)

	mutex.Unlock()
	wg.Done()
}
