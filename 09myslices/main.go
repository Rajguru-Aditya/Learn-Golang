package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices: ")

	var fruits = []string{"Apple", "Chikoo", "Mango"}
	fmt.Printf("type of fruits is %T\n", fruits)

	fruits = append(fruits, "Banana", "Peach")
	fmt.Println("fruits are", fruits)

	fruits = append(fruits[1:3])
	fmt.Println(fruits)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 2322
	highScores[2] = 2343
	highScores[3] = 22222
	highScores = append(highScores, 555, 666, 777)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// Remove value from slice based on index

	var courses = []string{"reactjs", "javascript", "python", "golang"}

	fmt.Println(courses)

	var index = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
