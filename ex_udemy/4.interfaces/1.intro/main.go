package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"example.com/note/note"
	"example.com/note/todo"
)


/****************************/
/*	INTERFACCIA: Creazione	*/			 				 
/****************************/	

/* 
* Interface Saver: Contiene la logica per salvare i dati
  Rappresenta un contratto: "Ogni struttura dati che implementa
  l'interfaccia Saver dovra implementare i suoi medoti"
*
* Metodi:	
	- Save() error
*/
type saver interface {
	Save() error
}



/********************************************************/
/*	INTERFACCIA: Funzione che fa uso del tipo "saver"	*/			 				 
/********************************************************/	
func saveData(data saver) error{

	//Con Save() = metodo dell'interfaccia, che avrà la logica per salvare i dati
	//             in base al tipo di struttura dati che passerò come argomento
	err := data.Save()

	if err != nil {
		fmt.Println("saving failed")
		return err
	}
	fmt.Println("saving successful")
	return nil
}




func getUserInput(prompt string) string{
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}


func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}







func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	//Creo un nuovo todo
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}


	todo.Display()

	/*********************************************************/
	/*	INTERFACCIA: Salvo i dati di "todo" usando il metodo */ 
	/*				 che fa uso dell'interfaccia	         */			 				 
	/*********************************************************/	
	err = saveData(todo)
	if err != nil {
		return
	}


	userNote.Display()

	/*************************************************************/
	/*	INTERFACCIA: Salvo i dati di "userNote" usando il metodo */ 
	/*				 che fa uso dell'interfaccia	             */			 				 
	/*************************************************************/	
	err = saveData(userNote)
	if err != nil {
		return
	}





}
