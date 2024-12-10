package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5


func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter number of years: ")
	fmt.Scan(&years)
	fmt.Println("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)


	// cosi assegno i valori di ritorno della funzione a due variabili (in ordine)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)
}


func calculateFutureValue(invAmount, expecRetRate, year float64) (float64, float64) {
	fv := invAmount * math.Pow(1 + expecRetRate/100, year)
	rfv := fv / math.Pow(1 + inflationRate/100, year)

	return fv, rfv
}