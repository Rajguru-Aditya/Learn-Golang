package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Hello, Welcome to function in golang")

	result := adder(3, 5)
	fmt.Println("The result is ", result)

	proResult, myMsg := proAdder(3, 5, 7, 9)
	fmt.Println("Pro result is ", proResult, " and ", myMsg)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi proAdder users"
}

func greeter() {
	fmt.Println("Namaste from golang")
}
