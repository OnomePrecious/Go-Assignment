package main

import (
	"fmt"
)

type Customer struct {
	AccountNumber    int
	BeginningBalance int
	Charges          int
	Credits          int
	CreditLimit      int
}

func main() {
	// Sample customer data
	customers := []Customer{
		{AccountNumber: 1001, BeginningBalance: 500, Charges: 200, Credits: 50, CreditLimit: 700},
		{AccountNumber: 1002, BeginningBalance: 1000, Charges: 300, Credits: 100, CreditLimit: 1500},
		{AccountNumber: 1003, BeginningBalance: 200, Charges: 1000, Credits: 200, CreditLimit: 1200},
	}

	// Check credit limit for each customer
	for _, customer := range customers {
		newBalance := customer.BeginningBalance + customer.Charges - customer.Credits
		fmt.Printf("Customer Account: %d\nNew Balance: %d\n", customer.AccountNumber, newBalance)
		if newBalance > customer.CreditLimit {
			fmt.Println("Credit limit exceeded")
		}
		fmt.Println("---------------------------")
	}
}
