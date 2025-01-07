package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"example.com/note/note"
	"example.com/note/todo"
)


/*********************/
/*	INTERFACCIA: Cre */
/*********************/	

type saver interface {
	Save() error
}

//Embedded interface: Ogni valore del ipo "outputtable" è anche un valore del tipo "saver"
//                    (avrà quindi il metodo Save()) ed avrà anche il metodo Display()
type outputtable interface {
	saver
	Display()
}




func saveData(data saver) error{

	err := data.Save()

	if err != nil {
		fmt.Println("saving failed")
		return err
	}
	fmt.Println("saving successful")
	return nil
}


/*Funzione che fa uso di saveData() e Display()
*
* Nota: La differenza qui è che io non posso usare il tipo "saver"
*		perchè non ha il metodo Display() e neanche il tipo "displayer"ù
*		perchè non ha il metodo Save()
*/
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
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
	
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}


	//uso outputData() per salvare i dati di todo
	err = outputData(todo)
	if err != nil {
		return
	}

	//uso outputData() per salvare i dati di userNote
	err = outputData(userNote)
	if err != nil {
		return
	}

}
