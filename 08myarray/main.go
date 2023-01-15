package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Peach"

	fmt.Println("Fruits are:", fruits)
	fmt.Println("Length of Fruits are:", len(fruits))

	var veggies = [3]string{"potatoes", "onions", "cabbage"}
	fmt.Println("Veggies are:", veggies)
	fmt.Println("Length Veggies are:", len(veggies))
}
