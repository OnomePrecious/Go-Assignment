package main

import "fmt"

const salary = 200
const commission = 0.09

func main() {
	var total float64

	for {
		var sales float64
		fmt.Print("Enter sales or 0 to finish: ")
		fmt.Scanln(&sales)

		if sales == 0 {
			break
		}

		total += sales
	}

	earnings := calculateEarnings(total)
	fmt.Printf("Total sales: $%.2f\n", total)
	fmt.Printf("Earnings: $%.2f\n", earnings)
}

func calculateEarnings(sales float64) float64 {
	return salary + sales*commission
}
