package main

import (
	"fmt"
	"math"
)

func main() {
	var number float64

	fmt.Println("N\tN2\tN3\tN4")
	fmt.Println()
	for i := 1; i <= 5; i++ {
		number++

		fmt.Println(number, "\t", math.Pow(number, 2), "\t", math.Pow(number, 3), "\t", math.Pow(number, 4))

	}
}
