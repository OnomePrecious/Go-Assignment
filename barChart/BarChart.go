package main

import (
	"fmt"
)

func main() {

	var numbers int
	var stars int
	fmt.Println("Enter five numbers between 1 and 30 : ")
	fmt.Scanln(&numbers)

	if numbers > 30 {
		fmt.Println("Numbers must be between 1 and 30 : ")
		return
	}

	for stars = 0; stars < numbers; stars++ {
		fmt.Print("*")
	}
}
