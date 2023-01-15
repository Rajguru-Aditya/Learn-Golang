package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// there is no inheritance in golang; no super or parent

	aditya := User{"Aditya", "aditya@go.dev", true, 20}
	fmt.Println(aditya)
	fmt.Printf("Aditya's details are: %+v\n", aditya)
	fmt.Printf("Name is %v and email is %v\n", aditya.Name, aditya.Email)
	aditya.GetStatus()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active: ", u.Status)
}
