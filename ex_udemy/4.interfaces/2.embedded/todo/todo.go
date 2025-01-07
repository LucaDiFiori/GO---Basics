
package todo

import (
	"errors"
	"fmt"
	"os"
	"encoding/json"
)


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



// Il tipo Todo implementa sia il metodo Display() che il metodo Save()
// rendendolo anche un tipo "outputtable"
func (todo Todo) Display() {
	fmt.Printf(todo.Text)
}

func (todo Todo) Save() error {
	
	fileName := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}