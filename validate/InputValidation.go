package main

import "fmt"

func main() {
	var count int
	var passes int
	var failures int
	var result int

	for i := 0; i < 10; i++ {
		fmt.Println("Enter result(1 = pass, 2 = fail): ")
		fmt.Scan(&result)
		if result != 1 && result != 2 {
			continue
		}
		if result == 1 {
			passes++

		} else {
			failures++
		}
		count = count + 1
		fmt.Printf("Passed: %d\nFailed: %d\n", passes, failures)

	}
}
