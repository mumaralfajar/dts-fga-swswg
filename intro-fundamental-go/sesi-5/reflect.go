package main

// import (
// 	"log"
// 	"reflect"
// )

// type UserSignup struct {
// 	Username string
// 	Password string
// }

// func handleSignupProcess(payload interface{}) {
// 	var reflectType reflect.Type = reflect.TypeOf(payload)

// 	if reflectType.Kind() != reflect.Struct {
// 		log.Panicf("type of payload has to be a struct, but what you gave is a %s \n", reflectType.Kind())
// 	}

// 	var reflectValue reflect.Value = reflect.ValueOf(payload)

// 	for i := 0; i < reflectType.NumField(); i++ {

// 		if reflectValue.Field(i).Interface() == "" {
// 			log.Panicf("%s is required \n", reflectType.Field(i).Name)
// 		}

// 	}
// }

// func main() {
// 	var user1 UserSignup = UserSignup{
// 		Username: "john@mail.com",
// 		Password: "123456",
// 	}

// 	var reflectType reflect.Type = reflect.TypeOf(user1)

// 	_ = reflectType

// fmt.Println("user1 field:", reflectType.NumField())

// fmt.Printf("reflectType.Field(0): %+v \n", reflectType.Field(0))

// fmt.Printf("reflectType.Field(0).Name: %s \n", reflectType.Field(0).Name)

// fmt.Println("reflectType.Kind() == reflect.Struct:", reflectType.Kind() == reflect.Struct)

// fmt.Println("reflectType.Kind():", reflectType.Kind())

// handleSignupProcess(user1)
// }
