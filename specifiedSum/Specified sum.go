package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	sum := 0
	for sum < number {
		var num int
		fmt.Print("Enter another number: ")
		fmt.Scan(&num)
		sum += num
	}

	fmt.Println("Sum of the integers is either greater than the input in the beginning or equal to the input.")
}
