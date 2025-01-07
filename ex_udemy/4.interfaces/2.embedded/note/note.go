package note

import (
	"errors"
	"fmt"
	"os"
	"time"
	"strings"
	"encoding/json"
)


type Note struct {
	Title		string		`json:"title"`
	Content 	string		`json:"content"`
	CreatedAt	time.Time	`json:"created_at"`
}


func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid value")
	}

	return Note {
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}


// Il tipo Todo implementa sia il metodo Display() che il metodo Save()
// rendendolo anche un tipo "outputtable"
func (n Note) Display() {
	fmt.Printf("Your note titled: %v has the following content: \n\n%v\n\n", n.Title, n.Content)
}

func (n Note) Save() error {
	
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}