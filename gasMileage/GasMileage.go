package main

import "fmt"

func main() {
	var count int
	var stop int
	var totalMiles float64
	var totalGallons float64
	var totalMilesPerGallon float64

	for {
		var milesDriven float64
		var gallonsUsed float64

		fmt.Printf("Enter the number of miles driven %d: ", count+1)
		fmt.Scanf("%f", &milesDriven)

		if milesDriven == -1 {
			break
		}
		fmt.Printf("Enter the number of gallons used %d: ", count+1)
		fmt.Scanf("%f", &gallonsUsed)

		totalMiles += milesDriven
		totalGallons += gallonsUsed
		totalMilesPerGallon += milesDriven / gallonsUsed
		count++

		fmt.Println("if you are tired enter -1 or press 1 to continue: ")
		fmt.Scanf("%d", &stop)
		if stop == -1 {
			break
		}
	}
	fmt.Printf("Total miles driven: %.2f\n", totalMiles)
	fmt.Printf("Total gallons used: %.2f\n", totalGallons)
	fmt.Printf("Total miles per gallon: %.2f\n", totalMilesPerGallon)
	if count > 0 {
		fmt.Printf("Average miles per gallon: %.2f\n", totalMilesPerGallon/float64(count))
	} else {
		fmt.Println("No trips in the record.")
	}
}
