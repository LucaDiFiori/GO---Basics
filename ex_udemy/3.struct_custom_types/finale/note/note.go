// Questo package contiente le informazioni della nota

package note

import (
	"errors"
	"fmt"
	"os"
	"time"
	"strings"
	"encoding/json"
)

/*
 * Note struct: Terrò "privati" i campi della struttura dati
 *              in modo che non possano essere modificati dall'esterno
 *				Renderò i campi accessibili tramite i metodi "Get" e "Set"
 *
 * - Aggiungo i tag json per rendere i campi della struct Note serializzabili
 */
type Note struct {
	Title		string		`json:"title"`
	Content 	string		`json:"content"`
	CreatedAt	time.Time	`json:"created_at"`
}


/*
 * Constructor: Inizializza la struttura dati Note
 * 
 * @Return:
 * - Note: struttura dati Note
 * - error: errore ed una Note vuota se il titolo o il contenuto sono vuoti 
*/
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


//Getters
func (n Note) Display() {
	fmt.Printf("Your note titled: %v has the following content: \n\n%v\n\n", n.Title, n.Content)
}


/* Salva i dati su un file chiamato come il titolo della nota
*
* - Sostituisce gli spazi con gli underscore con la funzione strings.ReplaceAll:
       - primo argomento: stringa da modificare
	   - secondo argomento: stringa da cercare
	   - terzo argomento: stringa con cui sostituire
*
* - Converte il titolo in minuscolo con la funzione strings.ToLower
*   e aggiunge l'estensione .json
*
* - formatto la struct note in formato JSON
	- NOTA: La funzione marshal può convertire solo i campi pubblici di una struct
*
* - return os.WriteFile(fileName, json, 0644):
*       - WriteFIle restituisce un errore. Se l'errore è nil, la funzione ha avuto successo
*         altrimenti restituisce l'errore
*/
func (n Note) Save() error {
	
	fileName := strings.ReplaceAll(n.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}