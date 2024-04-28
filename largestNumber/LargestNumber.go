package main

import "fmt"

func main() {
	var counter int
	var largest int

	for {
		var number int
		for counter = 0; counter < 10; counter++ {
			fmt.Println("Enter a number: ")
			_, err := fmt.Scanln(&number)

			if err != nil {
				fmt.Println("Error in input: ", err)
			} else {
				if number > largest {
					largest = number
				}
			}
			if number == 0 {
				break
			}
		}
		fmt.Printf("The largest number is: %d\n", largest)
	}
}
