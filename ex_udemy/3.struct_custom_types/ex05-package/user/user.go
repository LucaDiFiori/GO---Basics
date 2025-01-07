//Creiamo un package user che contiene un tipo User

//Modifica le iniziali maiuscole di user.go di ciò che voglio rendere
//pubblico (quindi visibile nel package main)

package user

import (
	"errors"
	"fmt"
	"time"
)

/*nota: Per rendere visibile sia il TIPO user che i suoi CAMPI,
*       tutto dovrà iniziare con la lettera maiuscola.
*
*       In questo esempio, avendo una constructor function newUser, terrò
*		terraò i campi privati e renderò pubblico solo il tipo User
*		L'inializzazione di un nuovo User avverrà tramite la funzione newUser
*/
type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

//Renderò il costruttore newUser pubblico 
func New(userFirstName, userLastName, userBirthdate string) (*User, error) {
	
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("all fields are required")
	}

	return &User {
		firstName: userFirstName,
		lastName: userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil
}

func (u *User) OutputUserDetails () {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = "Default"
	u.lastName = "Default"
}


