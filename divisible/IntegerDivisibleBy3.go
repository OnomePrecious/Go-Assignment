package main

import "fmt"

var sum int

func main() {
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
