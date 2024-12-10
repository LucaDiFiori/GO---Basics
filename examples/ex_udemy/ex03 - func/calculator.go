package main

import (
	"fmt"
)


func getUserINput(infoText string) float64{
	var infoInput float64
	fmt.Println(infoText)
	fmt.Scan(&infoInput)

	return infoInput
}


func calculate(revenue, expenses, taxRate float64) (inc, net, rat float64) {
	inc = revenue - expenses
	net = inc * (1 - taxRate/100)
	rat = inc / net

	return
}



func main() {
	fmt.Println("Welcome to the calculator")

	revenue := getUserINput("Enter Expected Revenue: ")
	expenses := getUserINput("Enter Expected Expenses: ")
	taxRate := getUserINput("Enter Expected Tax Rate: ")


	incomeBeforTax, netIncome, ratio := calculate(revenue, expenses, taxRate)


	fmt.Printf("Income before tax: %.2f\n", incomeBeforTax)
	fmt.Printf("Net income: %.2f\n", netIncome)
	fmt.Printf("Ratio: %.2f\n", ratio)
}



