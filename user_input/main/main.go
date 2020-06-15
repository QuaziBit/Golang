package main

import (
	"fmt"
)

func userInput() int {

	var num int

	fmt.Scan(&num)

	return num
}

func main() {
	var input int
	fmt.Printf("\n")
	for i := 0; i < 9; i++ {
		fmt.Printf("Enter number or -1 to stop: ")
		num := userInput()
		input = num
		fmt.Printf("You entered: %d\n", num)
		if num == -1 {
			break
		}
	}
	fmt.Printf("Last number you entered: %d\n", input)
}
