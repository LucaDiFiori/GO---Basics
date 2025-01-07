/* In questo esempio vediamo come utilizzare i puntatori per 
 * modificare il valore di una variabile
*/

package main

import "fmt"

func getAdultYears(age *int){

	//qui assegno alla posizione di memoria di age, il valore contenuto in age - 18
	//in questo modo la variabile age viene modificata
	*age = *age -18
}

func main() {
	age := 32

	fmt.Println("Age:", age)

	getAdultYears(&age)
	fmt.Println("Years since adult:", age)

	
}