package main

import (
	"fmt"
	"time"
)

//Nota: sia il nome della struct che i suoi campi iniziano con la lettera maiuscola
//      questo li rende "privati" e quindi non accessibili dall'esterno del package
type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

//passaggio di una variabile di una struct alla funzione
func outputUserDetails (u user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	//dichiarazione della variabile appUser di tipo user
	var appUser user

	// Assegno il valore alla variabile user. Questa sarà un'istanza della struct user
	// Questa è chiamata "Struct literal notation"
	appUser = user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	/*Ilizzando i campi della struct nell'ordine in cui sono stati definiti
	  posso anche scrivere:

	  	appUser = user{
			userFirstName,
			userLastName,
			userBirthdate,
			time.Now(),
	}
	*/
		
	//passaggio della variabile appUser alla funzione outputUserDetails
	outputUserDetails(appUser)

}
