package main

import "fmt"

func main() {

	//Con questa sintassi posso dichiarare un array
	// e poi inizializzarlo successivamente
	var productNames [4]string = [4]string{"A Book",}

	// O anche
	// var productNmaes [4]string
	// productNmaes = [4]string{"A Book"}

	//COn questa sintassi dovrò initializzare gli 
	//elementi dell'array al momento della dichiarazione
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}


	fmt.Println(prices)

	//stampo il terzo elemento dell'array prices
	fmt.Println(prices[2])

	//setto il terzo elemento dell'array productNames
	productNames[2] = "A Carpet"
	fmt.Println(productNames)

}