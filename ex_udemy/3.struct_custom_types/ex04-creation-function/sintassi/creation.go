//vediamo come creare una creation function

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

func (u user) outputUserDetails () {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

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



//Creation func: Inizializza i campi della struct e ritorna un valore di tipo user 
// con i campi inizializzati
func newUser(userFirstName, userLastName, userBirthdate string) user {
	return user{
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}
}


 
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	//inizializziamo la struct con la creation function
	appUser = newUser(userFirstName, userLastName, userBirthdate)


	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}