package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter username: ")

	// Comma ok || Err ok
	username, _ := reader.ReadString('\n')
	fmt.Println("Welcome ", username)
}
