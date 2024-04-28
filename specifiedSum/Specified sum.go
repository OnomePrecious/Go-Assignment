package main

import "fmt"

func main() {
	var target int
	fmt.Print("Enter a number: ")
	fmt.Scan(&target)

	sum := 0
	for sum < target {
		var num int
		fmt.Print("Enter an integer value: ")
		fmt.Scan(&num)
		sum += num
	}

	fmt.Println("Sum of the integers is either greater than the input in the beginning or equal to the input.")
}
