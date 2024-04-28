package main

import "fmt"

func main() {
	var firstNumber int
	fmt.Print("Enter first number: ")
	fmt.Scan(&firstNumber)

	var secondNumber int
	fmt.Print("Enter second number: ")
	fmt.Scan(&secondNumber)

	if firstNumber == secondNumber {
		fmt.Println(0)
	}
	if firstNumber > secondNumber {
		fmt.Println(1)
	}
	if secondNumber > firstNumber {
		fmt.Println(-1)
	}
}
