package main

import (
	"fmt"
	"example.come/struct/user"
)



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
 
func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")



	// Creo un nuovo User. Per farlo dovro usare questa sintassi
	// che signifca "appUser è un puntatore a User" che si 
	// trova nel package user 
	var appUser *user.User

	// Creo una nuova istanza di User (punta a un User)
	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}