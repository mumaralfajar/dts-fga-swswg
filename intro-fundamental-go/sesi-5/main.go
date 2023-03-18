package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	cpuCores := runtime.NumCPU()

	fmt.Println("cpuCores:", cpuCores)

	var wg sync.WaitGroup

	var mutex sync.Mutex

	fruits := []string{"Apple", "Mango", "Durian", "Grape", "Orange", "Pineapple"}

	result := []string{}

	wg.Add(len(fruits))

	for _, value := range fruits {

		go func(fruitItem string) {

			time.Sleep(2 * time.Second)

			mutex.Lock()

			result = append(result, fruitItem)

			mutex.Unlock()

			wg.Done()

		}(value)
	}
	wg.Wait()

	fmt.Printf("result: %#v \n", result)

	// wg.Add(len(fruits))

	// for _, f := range fruits {

	// 	go func(fruitItem string) {

	// 		mutex.Lock()

	// 		result = append(result, fruitItem)

	// 		mutex.Unlock()

	// 		wg.Done()

	// 	}(f)

	// }

}

func async(message string, wg *sync.WaitGroup) {
	fmt.Println(message)
	wg.Done()
}
