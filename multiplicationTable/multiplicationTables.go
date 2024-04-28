package main

import "fmt"

func main() {
	var number int
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println("Error in input: ", err)
	} else {

		for count := 1; count < 13; count++ {
			total := number * count
			fmt.Println(fmt.Sprintf("%d x %d = %d", number, count, total))
		}
	}

}
