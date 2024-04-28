package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	for {
		fmt.Print("Enter a five digit integer: ")
		fmt.Scanln(&input)

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number, please enter a five digit number.")
			continue
		}

		if num < 10000 || num > 99999 {
			fmt.Println("Invalid number, Please enter a five-digit integer.")
			continue
		}

		if isPalindrome(num) {
			fmt.Println(input, "is a palindrome.")
		} else {
			fmt.Println(input, "is not a palindrome.")
		}

		break
	}
}

func isPalindrome(num int) bool {
	input := strconv.Itoa(num)
	output := ""

	for i := len(input) - 1; i >= 0; i-- {
		output += string(input[i])
	}

	return input == output
}
