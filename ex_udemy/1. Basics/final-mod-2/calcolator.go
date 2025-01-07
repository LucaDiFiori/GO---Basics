package main

import (
	"fmt"
	"errors"
	"os"
)

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func storeResults(ebt, profit, ratio float64) {
	//Converto i valori in una singola stringa e li salvo in results
	results := fmt.Sprintf("Earnings Before Taxes: %2f\nNet Profit: %2f\nEBT/Net Profit Ratio: %2f\n", ebt, profit, ratio)

	//converto la stringa in bites
	data := []byte(results)

	os.WriteFile("calculations.txt", data, 0644)

}


func getUserInput(infoText string) (userInput float64, err error) {
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <=0 {
		return 0, errors.New("invalid input")
	}

	return userInput, nil
}



func main() {
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	storeResults(ebt, profit, ratio)
}
