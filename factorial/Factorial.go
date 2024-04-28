package main

import "fmt"

func main() {

	var count int
	var positiveNumber int
	var factorial int

	fmt.Println("Enter any positive number between 1 and 20:")
	fmt.Scan(&positiveNumber)

	if positiveNumber > 20 {
		fmt.Scan("Try again")

	} else {

		factorial = 1

		for count = 1; count <= positiveNumber; count++ {
			factorial *= count

			fmt.Println(positiveNumber, "===>", factorial)
		}
	}
}
