package main

import "fmt"

// variabili

/*
* NOTE:
* 1. Le variabili sono precedute dalla keyword "var", poi il nome della variaibile
*     poi il tipo di variabile
*
* 2. short variable declaration
* -----------------------------
* La dichiarazione abbreviata con := permette di dichiarare e inizializzare una 
* variabile in un unico passaggio senza dover specificare il tipo della variabile. 
* Go è un linguaggio tipizzato staticamente, ma grazie alla type inference 
* (inferenza del tipo), il compilatore è in grado di dedurre automaticamente il 
* tipo dalla parte destra dell'assegnazione.
*
* Questo tipo di dichiarazione può essere utilizzato SOLO ALL'INTERNO DI BLOCCHI,
* FUNZIONI o DICHIARAZIONI LOCALI. Non può essere usato per dichiarare variabili 
* globali fuori da una funzione.
*
* Con := devi sempre fornire un valore iniziale.
*
*/
func main() {

	// STRINGS__________________________________________________________________

	// metodo 1
	var nameOne string = "mario"

	// metodo 2
	var nameTwo = "luigi"

	// metodo 3
	var nameThree string
	nameThree = "piero"

	fmt.Println(nameOne, nameTwo, nameThree)

	//posso poi cmabiare il valore della variabile
	nameOne = "giovanni"
	fmt.Println(nameOne)

	// metodo 4: short variable declaration
	nameFour := "francesco"
	fmt.Println(nameFour)
	

	//INT, MEMORY & BEATS_______________________________________________________
	/*
	* Quanto detto sopra rimane valiso.
	* Nota su questi tipi: type 8 = intero a 8 bit. Abbiamo anche int16, int32, int64
	*/
	var numOne int8 = 127
	fmt.Println(numOne)

}