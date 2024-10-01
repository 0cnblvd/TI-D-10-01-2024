package main

import (
	"fmt"
)

func main() {

	const months = 12
	var expenses [months]float64

	//monthly expenses

	for i := 0; i < months; i++ {
		fmt.Printf("Enter the expenses for month %d: ", i+1)
		fmt.Scan(&expenses[i])
	}

	//perhitungan total
	var total float64
	for _, expense := range expenses {
		total += expense

	}
	average := total / months

	//output

	fmt.Printf("Total expenses for the year: %.2f\n", total)
	fmt.Printf("Average monthly expense: %.2f", average)

}
