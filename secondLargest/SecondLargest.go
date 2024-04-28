package main

import "fmt"

func main() {
	var counter int
	var largest int
	var secondLargest int

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
				} else if number > secondLargest {
					secondLargest = largest
				}
			}
		}
		fmt.Printf("The largest number is: %d\n", largest)
		fmt.Printf("The second largest number is: %d\n", secondLargest)
	}
}
