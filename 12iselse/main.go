package main

import "fmt"

func main() {
	fmt.Println("If else in GoLang")

	loginCount := 11
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Super User"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if num := 5; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is greater than 10")
	}

}
