package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func main() {
// 	defer catchError()

// 	if result, err := convertCharToInteger("10p"); err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("result:", result)
// 	}

// }

// func catchError() {
// 	if r := recover(); r != nil {
// 		fmt.Println("error:", r)
// 	} else {
// 		fmt.Println("Success")
// 	}
// }

// func convertCharToInteger(value string) (int, error) {
// 	num, err := strconv.Atoi(value)

// 	return num, err
// }
