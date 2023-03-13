package main

// import "fmt"

// type MyFunc func(text string, name string) string

// func main() {

// 	var g string = greet("John", "Halo")
// 	var sum1 = sum(20, 30, "John")
// 	if _, isValid := isValidText("Hal000"); isValid == true {
// 		// fmt.Println("Message:", text)
// 	} else {
// 		// panic("Message: " + text)
// 	}

// 	// ... => ellipsis
// 	// & => ampersand

// 	_, _ = g, sum1

// 	fruits := []string{"Apple", "Mango", "Durian", "Apple", "Durian", "Banana", "Mango"}

// 	fruitGroup, fruitAmount := handleFruitGroup(fruits)

// 	scores := []int{100, 40, 50, 60}

// 	result := sumScores("John", scores...)

// 	_, _ = result, scores

// 	fmt.Println("result:", result)

// 	fmt.Println("fruitGroup:", fruitGroup)
// 	fmt.Println("fruitAmount:", fruitAmount)

// 	_, _, _ = fruits, fruitGroup, fruitAmount

// 	var helloWorld MyFunc = func(text, name string) string {
// 		result := fmt.Sprintf("%s, %s", text, name)

// 		return result
// 	}

// 	var helloJohn MyFunc = func(text string, name string) string {
// 		return text + name
// 	}

// 	_ = helloJohn

// 	var resultOfHelloWorld = helloWorld("halo", "world")

// 	_ = resultOfHelloWorld

// 	var numbers []int = []int{23, 20, 10, 43, 18}

// 	var isOddNumber func(num int) bool = func(num int) bool {

// 		return num%2 != 0
// 	}

// 	var oddNumbers int = findOddNumbers(numbers, isOddNumber)

// 	_ = oddNumbers

// 	var selfInvocationFunc string = func() string {
// 		return "Hello World"
// 	}()

// 	fmt.Println("selfInvocationFunc:", selfInvocationFunc)
// }

// func findOddNumbers(numbers []int, callback func(int) bool) int {
// 	var result int
// 	for _, value := range numbers {
// 		if callback(value) == true {
// 			result++
// 		}
// 	}

// 	return result
// }

// func sumScores(name string, scores ...int) string {
// 	var result int

// 	for _, eachScore := range scores {
// 		result += eachScore
// 	}

// 	return fmt.Sprintf("Halo %s, total nilainya adalah %d yah :) \n", name, result)
// }

// func handleFruitGroup(fruits []string) (groupResult map[string]int, length int) {
// 	var result map[string]int = map[string]int{}

// 	for _, eachFruit := range fruits {
// 		result[eachFruit]++
// 	}

// 	groupResult = result

// 	length = len(fruits)

// 	return
// }

// func isValidText(text string) (string, bool) {
// 	var chars = []rune(text)

// 	if len(chars) < 5 {
// 		return "Text tidak valid", false
// 	}

// 	return "Text valid", true
// }

// func sum(num1, num2 int, name string) map[string]string {
// 	result := num1 + num2

// 	message := fmt.Sprintf("Halo %s, %d + %d sama dengan %d ya :)", name, num1, num2, result)

// 	response := map[string]string{
// 		"message": message,
// 	}

// 	return response
// }

// func greet(name, text string) string {
// 	return text + name
// }
