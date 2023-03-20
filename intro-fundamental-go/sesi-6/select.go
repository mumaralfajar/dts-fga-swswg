package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	c1 := make(chan string)

// 	c2 := make(chan string)

// 	go func() {
// 		time.Sleep(3 * time.Second)

// 		c1 <- "#1 Message"
// 	}()

// 	go func() {
// 		time.Sleep(2 * time.Second)

// 		c2 <- "#2 Message"
// 	}()

// 	for i := 1; i <= 2; i++ {
// 		select {
// 		case result1 := <-c1:
// 			fmt.Println("received from c1:", result1)
// 		case result2 := <-c2:
// 			fmt.Println("received from c2:", result2)
// 		}
// 	}

// }
