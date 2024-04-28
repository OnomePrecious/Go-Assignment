package main

import "fmt"

func main() {
	var baseLength int

	fmt.Print("Enter the length of the base of a triangle between 1-10: ")
	fmt.Scanln(&baseLength)

	if baseLength < 1 || baseLength > 10 {
		fmt.Println("Invalid length, the length should be between 1 and 10.")
		return
	}

	for i := 1; i <= baseLength; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
