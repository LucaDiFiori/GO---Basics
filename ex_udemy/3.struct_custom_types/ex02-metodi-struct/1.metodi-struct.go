package main

import (
	"fmt"
	"time"
)


type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

/*Posso creare un metodo per la struct user in questo modo:
* 1. Creo una funzione con il nome del metodo che voglio creare
* 2. Prima del nome della funzione, inserisco il nome della struct a cui il metodo appartiene 
     fra parentesi tonde

* Non serve passare la variabile della struct come argomento, 
* perchè il metodo è già associato alla struct
*/
func (u user) outputUserDetails () {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}




func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
 
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")


	appUser := user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// Chiamo il metodo outputUserDetails() per la struct user come
	// se fosse un suo campo
	appUser.outputUserDetails()

}