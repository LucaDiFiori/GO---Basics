// Supponiamo di voler creare una slice del nostro array
// `prices` che contenta gli elementi centrali (1 e 2)

package main

import "fmt"

func main() {
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

	//dichiaro la slice `featuredPrices` basata sull'array `prices`
	//che conterrà gli elementi prices[1] e prices[2]
	featuredPrices := prices[1:3]

	fmt.Println(featuredPrices)
}