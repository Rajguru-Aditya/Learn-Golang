package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// there is no inheritance in golang; no super or parent

	aditya := User{"Aditya", "aditya@go.dev", true, 20}
	fmt.Println(aditya)
	fmt.Printf("Aditya's details are: %+v\n", aditya)
	fmt.Printf("Name is %v and email is %v", aditya.Name, aditya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
