package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	loops := 5

// 	ch := make(chan string)

// 	wg := sync.WaitGroup{}

// 	wg.Add(loops)

// 	for i := 1; i <= loops; i++ {
// 		go func(num int) {
// 			message := fmt.Sprintf("Message #%d", num)

// 			fmt.Printf("Start Sending Data #%d \n", num)

// 			ch <- message

// 			wg.Done()

// 		}(i)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(ch)
// 	}()

// 	for data := range ch {
// 		fmt.Println("recieve data:", data)
// 	}

// go func() {

// }()

// go func() {
// 	for i := 1; i <= loops; i++ {
// 		message := fmt.Sprintf("Message #%d", i)
// 		fmt.Printf("Start Sending Data #%d \n", i)

// 		ch <- message

// 		fmt.Printf("After Sending Data #%d \n", i)

// 	}
// 	close(ch)
// }()

// time.Sleep(2 * time.Second)

// for data := range ch {

// 	fmt.Println("recieve data:", data)

// 	time.Sleep(2 * time.Second)
// }

// for i := 1; i <= loops; i++ {
// 	time.Sleep(2 * time.Second)

// 	result := <-ch
// 	fmt.Println("recieve data:", result)
// }

// }
