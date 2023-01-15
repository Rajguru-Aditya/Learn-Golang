package main

import "fmt"

func main() {
	defer fmt.Print("Defer in Golang")
	defer fmt.Println("World")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
