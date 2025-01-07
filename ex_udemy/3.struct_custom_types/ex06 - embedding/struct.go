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


	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}


	//Creo un nuovo Admin
	admin := user.NewAdmin("test@example.come", "password")

	//Uso i metodi della struct User
	admin.OutputUserDetails()

	/*  Se avessi assegnato un nome al campo User nella struct Admin,
		avrei dovuto usare quel nome per accedere ai metodi e campi
		della struct User:

		admin.u.OutputUserDetails()
	*/


	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()

}