package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	fmt.Println(days)

	// for loop
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("index is %v and value is %v\n", index, day)
	// }

	anyValue := 1
	for anyValue < 10 {

		if anyValue == 5 {
			goto myLabel
		}

		fmt.Println("value is: ", anyValue)
		anyValue++
	}

myLabel:
	fmt.Println("I am outside the loop")

}
