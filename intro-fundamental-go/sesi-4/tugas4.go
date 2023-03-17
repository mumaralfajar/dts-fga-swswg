package main

import (
	"errors"
	"fmt"
	"os"
)

type Student struct {
	Id        string
	Name      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var students = []Student{
	{
		Id:        "go-1",
		Name:      "John",
		Alamat:    "Jakarta",
		Pekerjaan: "Programmer",
		Alasan:    "Mau belajar golang",
	},
	{
		Id:        "go-2",
		Name:      "Jane",
		Alamat:    "Bandung",
		Pekerjaan: "Programmer",
		Alasan:    "Mau belajar golang",
	},
	{
		Id:        "go-3",
		Name:      "Jenny",
		Alamat:    "Bogor",
		Pekerjaan: "Programmer",
		Alasan:    "Mau belajar golang",
	},
	{
		Id:        "go-4",
		Name:      "Jen",
		Alamat:    "Bekasi",
		Pekerjaan: "Programmer",
		Alasan:    "Mau belajar golang",
	},
}

func GetStudentById(id string) (*Student, error) {
	for _, student := range students {
		if student.Id == id {
			return &student, nil
		}
	}

	return nil, errors.New("student not found")
}

func main() {
	var args = os.Args

	if len(args) < 2 {
		fmt.Println("Please provide student id")
		return
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	var student, err = GetStudentById(args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v \n", student)
}
