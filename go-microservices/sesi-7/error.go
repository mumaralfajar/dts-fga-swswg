package main

// import "fmt"

// type CustomError string

// func (ce CustomError) Error() string {
// 	fmt.Println("Invoked()")
// 	return string(ce)
// }

// func main() {

// 	if result, err := isValidName("Agus"); err != nil {
// 		fmt.Println("error:", err)
// 	} else {
// 		fmt.Println("result:", result)
// 	}
// }

// func isValidName(name string) (string, error) {
// 	if len([]rune(name)) < 5 {
// 		return "", CustomError("invalid name")
// 	}

// 	return name, nil
// }
