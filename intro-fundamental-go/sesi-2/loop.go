package main

import "fmt"

func main() {

	// var j = 0

	// for j < 5 {
	// 	fmt.Println("j:", j)
	// 	j++
	// }

	var j = 0

	// for {
	// 	if j == 5 {
	// 		break
	// 	}

	// 	fmt.Println("j:", j)
	// 	j++
	// }

	var numbers []int = []int{
		100,
		200,
		300,
		400,
		500,
	}

	_, _ = numbers, j

	// for index, value := range numbers {
	// 	fmt.Println("index:", index, "value:", value)
	// }

	var city string = "Jakarta"

	_ = city

	var russianCity string = "Город"

	for index, value := range russianCity {
		fmt.Printf("character %#U starts at byte position %d \n", value, index)
	}

	// var charsOfRussianCity []rune = []rune(russianCity)

	// fmt.Println(len(charsOfRussianCity))

	// for i := 0; i < len(charsOfRussianCity); i++ {
	// 	fmt.Println(string(charsOfRussianCity[i]))
	// }

	// fmt.Println(len(russianCity))

	// for i := 0; i < len(russianCity); i++ {
	// 	fmt.Println(string(russianCity[i]))
	// }

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(numbers[i])
	// }
}
