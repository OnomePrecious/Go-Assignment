package main

import "fmt"

// FOR STATEMENTS

//func main() {
//	i := 1
//	for i <= 10 {
//		fmt.Println(i)
//		i = i + 1
// the i = i + 1 line is extremely important, because without it i <= 10 would always evaluate to true and our program would never stop.

// }
// }
// shorter way of doing the loop
//func m

// IF STATEMENTS

//func main() {
//	for i := 1; i <= 10; i++ {
//		if i%2 == 0 {
//			fmt.Println(i, "is an even number")
//		} else {
//			fmt.Println(i, "is an odd number")
//		}
//	}
//}

// SWITCH STATEMENT
func main() {
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	switch input {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}
}
