package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"example.com/note/note"
)

/*Funzione per stampare il prompt e ricevere l'input dall'utente
* tramite bufio.NewReader(os.Stdin)
* 
* @Return:
* - string: valore inserito dall'utente
* - error: errore se l'utente non inserisce nulla 
*/
func getUserInput(prompt string) string{
	fmt.Printf("%v ", prompt)

	//salvo il reader in una variabile
	reader := bufio.NewReader(os.Stdin)

	// uso la variabile reader per leggere la stringa inserita dall'utente
	// Nota: ReadString('\n') legge fino a quando non trova un carattere specificato
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	//La stringa estratta conterrà lo \n (o \r) finale, quindi la elimino
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}



/*Funzione per ottenere i titolo e contenuto della nota
*
* @Return:
* - string: titolo della nota
* - string: contenuto della nota
* - error: errore se l'utente non inserisce nulla o se c'è un errore nella funzione getUserInput
*
* Nota: getNoteData, in caso di errore, "progagherà" il messaggio di errore
*       della funzione getUserInput al main
*/
func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}



/*
* Funzione principale
* - Ottiene il titolo e il contenuto della nota con la funzione getNoteData
* - Crea una nuova nota con il costruttore New
* - Stampa la nota con il metodo Display
* - Salva la nota con il metodo Save
*/
func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()

	//NOta: La variabile err già esiste, quindi non serve usare := ma solo =
	err = userNote.Save()
	if err != nil {
		fmt.Println("saving the note failded")
		return
	}

	fmt.Println("Note saved successfully")

}
