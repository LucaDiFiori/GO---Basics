/* In questo esempio vediamo come utilizzare i puntatori per 
 * passare un valore per riferimento ad una funzione evitando 
 * di creare una copia della variabile originale.
*/

package main

import "fmt"

func getAdultYears(age *int) int {
	return *age - 18
}

func main() {
	age := 32

	// agePtr contains the memory address of the variable age
	agePtr := &age

	fmt.Println("Age:", age)
	fmt.Println("AgePtr:", agePtr)
	fmt.Println("AgePtr value:", *agePtr)

	adultYears := getAdultYears(&age)
	fmt.Println("Years since adult:", adultYears)
	
}