package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}


	featuredPrices := prices[1:]

	//creo una slice dalla slice featuredPrices
	highlightPrices := featuredPrices[:1]

	fmt.Println(highlightPrices)
}