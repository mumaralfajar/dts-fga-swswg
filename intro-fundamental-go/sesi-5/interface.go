package main

// import "fmt"

// type UserHandler interface {
// 	SignUp()

// 	SignIn()
// }

// type Customer struct {
// 	Name string
// }

// func (c Customer) SignIn() {
// 	fmt.Println("customer SignIn()")
// }

// func (c Customer) SignUp() {
// 	fmt.Println("customer SignUp()")
// }

// func (c Customer) TopUp() {
// 	fmt.Println("customer TopUp()")
// }

// type MyString string

// func (m MyString) SignIn() {
// 	fmt.Println("MyString  SignIn()")
// }

// func (m MyString) SignUp() {
// 	fmt.Println("MyString  SignUp()")
// }

// type Product struct {
// 	Title string
// }

// func (p *Product) SignIn() {
// 	fmt.Println("Product SignIn()")
// }

// func (p *Product) SignUp() {
// 	fmt.Println("Product SignUp()")
// }

// func main() {

// 	var p UserHandler = &Product{}

// 	_ = p

// 	var uh UserHandler = Customer{}

// 	uh.SignIn()

// 	uh.SignUp()

// 	uh = MyString("Hai")

// 	var result = uh.(Customer)

// 	fmt.Println(result)

// 	var randomData interface{} = 20

// 	randomData = "John"

// 	randomData = true

// 	randomData = map[string]string{}

// 	var isValid, _ bool = randomData.(bool)

// 	_ = isValid
// }
