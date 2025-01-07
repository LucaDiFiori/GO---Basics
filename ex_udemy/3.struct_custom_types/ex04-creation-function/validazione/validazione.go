

//vediamo come validare i dati inseriti dall'utente 
// attraverso una funzione di validazione

package main

import (
	"fmt"
	"time"
	"errors"
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
	fmt.Scanln(&value)
	return value
}



//Aggiungiamo delle validazioni per i dati inseriti dall'utente
func newUser(userFirstName, userLastName, userBirthdate string) (*user, error) {
	
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &user {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}


 
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")



	var appUser *user

	//NOTA: Per non dichiarare "err" posso usare ":=". Questo dichiarerà
	//      solo la variabile "err" e la inizializzerà con il valore restituito dalla funzione
	//      senza dichiarare la variabile "appUser"
	appUser, err := newUser(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}



	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()

}