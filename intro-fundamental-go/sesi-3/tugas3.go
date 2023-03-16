package main

import "fmt"

func main()  {
	greet := "selamat malam"

	for _, char := range greet {
		fmt.Printf("%c\n", char)
	}

	charCount := make(map[string]int)

	for _, char := range greet {
		charCount[string(char)]++
	}

	fmt.Println(charCount)
}
