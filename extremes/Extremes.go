package main

import "fmt"

func main() {
	var userInput int

	fmt.Print("Enter the number of integers: ")
	fmt.Scanln(&userInput)

	if userInput <= 0 {
		fmt.Println("Invalid input, enter a positive number")
		return
	}
	var minimumValue int
	var maximumValue int
	var sum int
	for i := 0; i < userInput; i++ {
		var value int
		fmt.Println("Enter the numbers: ")
		fmt.Scanln(&value)

		if i == 0 {
			minimumValue = value
			maximumValue = value
		} else {
			if value < minimumValue {
				minimumValue = value
			}
			if value > maximumValue {
				maximumValue = value
			}
		}

		sum += value
	}
	fmt.Println("The minimum value is: ", minimumValue)
	fmt.Println("The maximum value is: ", maximumValue)
	fmt.Println("The sum of the minimum and maximum value is: ", minimumValue+maximumValue)

}
