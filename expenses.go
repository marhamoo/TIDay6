package main

import (
	"fmt"
)

func main() {
	const months = 12
	var expenses [months]float64

	// Input monthly expenses
	for i := 0; i < months; i++ {
		fmt.Printf("Enter the expense for months %d: ", i+1)
		fmt.Scan(&expenses[1])
	}
	// Calculate total and average expenses
	var total float64
	for _, expense := range expenses {
			total += expense 
	}

	average := total / months
	fmt.Printf("Total expenses for the year: %.2f\n", total)
	fmt.Printf("Average monthly expense: %.2f\n", average)
}