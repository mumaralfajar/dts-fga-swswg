package main

import "fmt"

func main()  {
	for  i := 0; i <= 4; i++ {
		println("Nilai i = ", i)
	}

	for j := 0; j <= 10; j++ {
		if j == 5 {
			input := "САШАРВО"
			for index, char := range input {
				fmt.Printf("character %U '%c' starts at byte position %d\n", char, char, index)
			}
			continue
		}
		println("Nilai j = ", j)
	}
}
