package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	const inflationRate = 2.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Println("Enter number of years: ")
	fmt.Scan(&years)
	fmt.Println("Enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)


	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate/100, years) 
	futureRealValue := futureValue / math.Pow(1 + inflationRate/100, years)

	fmt.Println("Future value is", futureValue)
	fmt.Println("Future real value is", futureRealValue)
}
