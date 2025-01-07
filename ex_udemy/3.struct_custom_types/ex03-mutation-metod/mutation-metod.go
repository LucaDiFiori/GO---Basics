//vediamo come creare un metodo che modifica i campi di una struttura

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

//metodo 1: non modifica la struttura
func (u user) outputUserDetails () {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

//metodo 2: modifica la struttura
func (u *user) clearUserName() {
	u.firstName = "Default"
	u.lastName = "Default"
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

	appUser.outputUserDetails()

	//cancelletà il nome e il cognome nella struct
	appUser.clearUserName()
	appUser.outputUserDetails()

}