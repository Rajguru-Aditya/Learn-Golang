package main

import "fmt"

const LoginToken = "akhfbgbv" //Public variable

func main() {
	var username string = "Aditya"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var emptyString string
	fmt.Println(emptyString)
	fmt.Printf("Variable is of type: %T \n", emptyString)

	// implicit type
	var randomVar = 3000
	fmt.Println(randomVar)

	// without var --- only allowed inside a function
	users := 3000
	fmt.Println(users)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
