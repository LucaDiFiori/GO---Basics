//Vediamo "l'ereditarietà" fra struct in Go

package user

import (
	"errors"
	"fmt"
	"time"
)


type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// Creiamo un'altra classe che eredita da User
// Questa ereiterà tutti i campi e i metodi di User ed in più
// avrà i campi email e pasword
type Admin struct {
	email string
	password string

	//Uso semplicemente il nome della struct che voglio ereditare
	User //Posso anche dare un nome a questo campo es: u User
}

//Posso aggiungere metodi alla struct Admin. Ad esempio un costruttore
func NewAdmin (email, password string) Admin {
	return Admin {
		email: email,
		password: password,

		//Inizializzo la struct User. Per semplicità assegnamo direttamente dei nomi
		User : User {
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}

/*
Se avessi assegnato un nome al campo User nella struct Admin, 
avrei dovuto usare quel nome per accedere ai metodi e campi 
della struct User

Esempio:
func NewAdmin (email, password string) Admin {
	(...)
		u : User {
			firstName: "Admin",
			lastName: "Admin",
			birthdate: "01/01/2000",
			createdAt: time.Now(),
		},
	}
}
*/




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
	u.firstName = "___"
	u.lastName = "___"
}


