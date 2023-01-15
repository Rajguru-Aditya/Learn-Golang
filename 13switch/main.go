package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch case in GoLang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Random dice number", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You got 1 and you can start")
		fallthrough
	case 2:
		fmt.Println("You got 2 and you can move 2 steps")
	case 3:
		fmt.Println("You got 3 and you can move 3 steps")
	case 4:
		fmt.Println("You got 4 and you can move 4 steps")
	case 5:
		fmt.Println("You got 5 and you can move 5 steps")
	case 6:
		fmt.Println("You got 6 and you can move 6 steps and roll the dice again!")
	default:
		fmt.Println("What was that!")
	}

}
