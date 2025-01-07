package main

import
(
	"fmt"
)

func main() {

	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Println("Welcome to the calculator")

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	incomeBeforTax := revenue - expenses
	netIncome := incomeBeforTax * (1 - taxRate/100)

	fmt.Printf("Income before tax: %.2f\n", incomeBeforTax)
	fmt.Printf("Net income: %.2f\n", netIncome)
	
	ratio := netIncome / revenue
	fmt.Printf("Ratio: %.2f\n", ratio)
}