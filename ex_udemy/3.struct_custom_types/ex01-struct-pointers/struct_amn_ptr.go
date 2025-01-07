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

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}


func outputUserDetails (u *user) {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
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

	//per non eseguire una copia della variabile appUser, passiamo il puntatore alla variabile
	outputUserDetails(&appUser)

}