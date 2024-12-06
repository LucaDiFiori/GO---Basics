In Go (Golang), **la concorrenza** si riferisce alla capacità di un programma di eseguire più attività indipendenti in parallelo o quasi in parallelo, sfruttando al meglio le risorse disponibili. È una caratteristica fondamentale del linguaggio, progettata per semplificare la scrittura di codice concorrente.

### Concetti chiave della concorrenza in Go:

### Goroutine:
- Sono leggere unità di esecuzione gestite dal runtime di Go.
-  Sono simili ai thread, ma molto più leggere e meno costose in termini di risorse.
- Si avviano utilizzando la parola chiave `go` seguita da una funzione o chiamata di funzione.

 - Esempio:
 ```go
 package main

import (
    "fmt"
    "time"
)

func sayHello() {
    fmt.Println("Hello!")
}

func main() {
    go sayHello() // Avvia una goroutine
    fmt.Println("Main Function")
    time.Sleep(time.Second) // Attendi per vedere l'output
}
 ```
 
In questo esempio, la goroutine `sayHello` viene eseguita indipendentemente dal flusso principale.



***
### Channel
- Sono strutture di comunicazione sicure per lo scambio di dati tra goroutine.
- Consentono alle goroutine di sincronizzarsi e scambiarsi informazioni in modo sicuro.
- Si creano con la funzione `make`:

```go
ch := make(chan int) // Canale per scambiare numeri interi
```


- Esempio di invio e ricezione:
```go
package main

import "fmt"

func sendMessage(ch chan string) {
    ch <- "Hello from Goroutine" // Invio di un messaggio nel canale
}

func main() {
    ch := make(chan string)
    go sendMessage(ch)
    msg := <-ch // Ricezione di un messaggio dal canale
    fmt.Println(msg)
}
```


***
### **Sincronizzazione**:
- Grazie ai channel, Go offre un modo semplice e sicuro per sincronizzare le goroutine.
 - Il runtime di Go garantisce che non ci siano race condition se i channel vengono usati correttamente.


***
### Select:
- Una struttura simile a uno switch per gestire più operazioni sui channel.

Esempio:
```go
package main

import "fmt"

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() { ch1 <- "Message from ch1" }()
    go func() { ch2 <- "Message from ch2" }()
    
    select {
    case msg1 := <-ch1:
        fmt.Println(msg1)
    case msg2 := <-ch2:
        fmt.Println(msg2)
    }
}
```



***

### Differenza tra **concorrenza** e **parallelismo**:

- **Concorrenza**: più attività vengono gestite in modo indipendente, ma non necessariamente eseguite allo stesso tempo.
- **Parallelismo**: più attività vengono eseguite simultaneamente su processori diversi.

Go supporta entrambi i paradigmi, ma la concorrenza è il focus principale del linguaggio, rendendo semplice scrivere codice che sfrutti bene i moderni processori multi-core.
