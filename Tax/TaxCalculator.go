package main

import "fmt"

const taxRate = 0.15
const excessCeiling = 0.02
const earningsInUsd = 30000

func main() {
	citizen1 := getCitizenInfo(1)
	citizen2 := getCitizenInfo(2)
	citizen3 := getCitizenInfo(3)

	fmt.Printf("Total tax for %s: $%.2f\n", citizen1.name, calculateTax(citizen1.earnings))
	fmt.Printf("Total tax for %s: $%.2f\n", citizen2.name, calculateTax(citizen2.earnings))
	fmt.Printf("Total tax for %s: $%.2f\n", citizen3.name, calculateTax(citizen3.earnings))
}

type citizen struct {
	name     string
	earnings float64
}

func getCitizenInfo(citizens int) citizen {
	var name string
	var earnings float64
	fmt.Printf("Enter name of Citizen %d: ", citizens)
	fmt.Scanln(&name)
	fmt.Printf("Enter earnings of Citizen  %d: ", name)
	fmt.Scanln(&earnings)
	return citizen{name: name, earnings: earnings}
}

func calculateTax(earnings float64) float64 {
	var totalTax float64

	if earnings <= earningsInUsd {
		totalTax = earnings * taxRate
	} else {
		totalTax = earningsInUsd*taxRate + (earnings-earningsInUsd)*excessCeiling
	}

	return totalTax
}
