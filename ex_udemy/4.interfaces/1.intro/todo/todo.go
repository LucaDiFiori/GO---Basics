//

package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)

/*
* Todo strucr, Campi:
	- Text: testo della nota
*/
type Todo struct {
	Text string	`json:"text"`
}


func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("invalid value")
	}

	return Todo {
		Text: content,
	}, nil
}


func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}





/************************************************************************/
/*	INTERFACCIA: Implementazione del metodo save() per la struct todo	*/			 				 
/************************************************************************/	
func (todo Todo) Save() error {
	
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}