# GO---Basics

# Table of contents
- [](#)

***
***
# Panoramica di Go

Go (o Golang) è un linguaggio di programmazione open-source sviluppato da Google nel 2007, progettato per essere efficiente, conciso e facile da usare. È particolarmente apprezzato per la sua gestione della concorrenza, semplicità di sintassi, velocità di esecuzione e capacità di compilare programmi in binari eseguibili autonomi.

## 1. Semplicità e chiarezza
Go è stato progettato con l’obiettivo di mantenere il codice semplice e leggibile. La sintassi è minimalista e diretta, eliminando molte delle complessità presenti in altri linguaggi come C++ o Java.

- **Tipizzazione statica**: Go è fortemente tipizzato, ma il compilatore è in grado di inferire il tipo in molti casi.
- **Design ridotto**: Non ci sono ereditarietà, classi o costrutti complessi come generici (almeno fino alla versione 1.18), ma si punta molto sulla composizione e le interfacce.

## 2. Compilato e veloce
Go è un linguaggio compilato, il che significa che i programmi vengono tradotti direttamente in codice macchina, offrendo prestazioni simili a C o C++. Tuttavia, il tempo di compilazione è molto rapido, il che rende il ciclo di sviluppo più produttivo.

## 3. Concorrenza nativa
Go ha un supporto nativo alla concorrenza, basato su **goroutine** e **canali**. 
- Goroutine: Sono funzioni che possono essere eseguite in parallelo. Sono leggere e richiedono pochissima memoria (circa 2 KB per goroutine), il che permette di eseguire migliaia di goroutine in un programma.
```go
go funzione()  // Lancia una goroutine
```
- Canali: Forniscono un meccanismo sicuro per la comunicazione tra goroutine. Consentono di sincronizzare le goroutine e passare dati tra di loro.
```go
ch := make(chan int)  // Crea un canale di interi
```
Il modello di concorrenza di Go è ispirato dal modello CSP (Communicating Sequential Processes), che favorisce un approccio "non condiviso" alla concorrenza, rendendo più facile evitare errori come le race condition.

## 4. Garbage Collection
Go è dotato di un garbage collector integrato che gestisce automaticamente l'allocazione e la deallocazione della memoria. Ciò semplifica la gestione della memoria rispetto a linguaggi come C o C++, dove la memoria deve essere allocata e liberata manualmente. Tuttavia, il garbage collector è stato progettato per essere molto veloce e con latenza bassa, minimizzando l'impatto sulle prestazioni.

## 5. Struttura del codice
Go organizza il codice in package. Ogni file sorgente Go deve appartenere a un package, e ogni progetto Go inizia con un file main.go che appartiene al package main. Go ha una libreria standard molto ricca, e i package aiutano a modularizzare e riutilizzare il codice.

- **Package main**: È il punto di ingresso del programma e contiene la funzione main() che viene eseguita quando il programma parte.
- **Package standard**: Fornisce molte funzionalità integrate, come la gestione di input/output (fmt), il networking (net/http), la manipolazione di stringhe (strings), e così via.

## 6. Compilazione multipiattaforma
Go è progettato per essere facilmente compilato su diverse piattaforme. Puoi scrivere codice Go una volta e compilarlo per varie piattaforme (Windows, Linux, macOS, ecc.) senza modifiche. Ad esempio, puoi creare un eseguibile per un sistema operativo diverso usando il comando:

```go
GOOS=linux GOARCH=amd64 go build
```

## 7. Strumenti integrati
Go include una serie di strumenti integrati che aiutano nello sviluppo del codice:

- **go build**: Compila il codice Go.
- **go run**: Esegue il codice senza creare un binario.
- **go test**: Esegue test automatici sui package Go.
- **go fmt**: Formatta il codice Go secondo le convenzioni standard del linguaggio.
- **go doc**: Mostra la documentazione di package e funzioni.
- **go mod**: Gestisce le dipendenze tramite moduli Go (introdotto nella versione 1.11).

## 8. Interfacce
Le interfacce in Go sono uno dei principali meccanismi di astrazione. Le interfacce definiscono un insieme di metodi che un tipo deve implementare, ma a differenza di altri linguaggi come Java, non richiedono che il tipo esplicitamente "dichiari" che implementa un'interfaccia.

Esempio di interfaccia:

```go
type Shape interface {
    Area() float64
}
```
Qualsiasi tipo che implementa il metodo **Area()** automaticamente soddisfa l'interfaccia Shape, senza necessità di un’esplicita dichiarazione.


## 9. Error handling
Go non utilizza eccezioni, ma si affida alla gestione esplicita degli errori. La convenzione è che le funzioni che possono fallire restituiscano un valore di tipo **error** come secondo valore di ritorno, che deve essere gestito dal chiamante.

Esempio:

```go
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
```

## 10. Moduli e gestione delle dipendenze
Go gestisce le dipendenze attraverso il sistema di moduli, introdotto con Go 1.11. I moduli permettono di specificare versioni esatte delle dipendenze e di risolvere facilmente i conflitti tra versioni diverse. I file `go.mod` e `go.sum` vengono utilizzati per definire le dipendenze del progetto e garantirne la riproducibilità.

## 11. Ecosistema e utilizzi comuni
Go è ampiamente utilizzato per costruire:

- Sistemi distribuiti e microservizi (per esempio Docker e Kubernetes sono scritti in Go).
- Applicazioni di rete e server web ad alte prestazioni (grazie al supporto nativo per la concorrenza e il networking).
- Tool di sistema e strumenti di sviluppo.

***
***

# PACKAGE
In Go, un **package** è un'unità fondamentale di organizzazione del codice. Un package può contenere uno o più file sorgente Go che sono organizzati logicamente per fornire funzionalità correlate.

# Caratteristiche principali dei package in Go:
1. **Nome del package**: Ogni file sorgente Go inizia con una dichiarazione del package:
```go
package main
```
Il nome del package indica a quale insieme logico di codice appartiene il file. Ad esempio, un programma eseguibile usa package main, mentre i pacchetti di librerie personalizzate possono usare un nome specifico del dominio che descrive la loro funzionalità (come package math).

2. **Package main**: È speciale in Go. Quando un file sorgente appartiene al package main, significa che contiene il codice per un programma eseguibile. Deve includere una funzione main() che è il punto di ingresso del programma:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

3. **Importazione di package**: Per riutilizzare codice da altri package (sia built-in che creati da te), si usa la direttiva import. Ad esempio, per importare il package fmt per la stampa formattata:
```go
import "fmt"
```
Go viene fornito con un set di package standard (fmt, math, net/http, ecc.) che forniscono molte funzionalità di base. Puoi anche creare e importare pacchetti personalizzati.

4. **Visibilità (export):**: In Go, la visibilità delle funzioni, delle variabili o delle strutture è determinata dal fatto che il nome inizia con una lettera maiuscola o minuscola:
- Se il nome di una funzione o di una variabile inizia con una **lettera maiuscola**, essa è **esportata** (visibile e utilizzabile da altri package).
- Se il nome inizia con una lettera **minuscola**, è **non esportata** (visibile solo all'interno del package).
```go
package math

// Esportata (pubblica)
func Add(a int, b int) int {
    return a + b
}

// Non esportata (privata)
func subtract(a int, b int) int {
    return a - b
}
```

5. **Organizzazione dei file**: I package possono essere divisi in più file, ma tutti i file devono dichiarare lo stesso nome di package in cima. Ad esempio, tutti i file che fanno parte del package utils devono iniziare con:
```go
package utils
```

6.**Directory**: In Go, la struttura delle directory riflette i package. Ogni directory è un package, quindi i file sorgente all'interno di una directory devono appartenere allo stesso package.

## Esempio
Creiamo un programma che usa un package personalizzato.
Nella directory greeting, creiamo i file che costituiranno questo package:
Supponiamo che il package greeting sia composto da due file: greet.go e farewell.go. Tutti i file risiederanno nella stessa directory e avranno il nome del package dichiarato come greeting.
**File grert.go**
```go
// greet.go
package greeting

import "fmt"

// Funzione esportata (pubblica)
func Hello() {
    fmt.Println("Hello!")
}
```

**file farewell.go**
```go
// farewell.go
package greeting

import "fmt"

// Funzione esportata (pubblica)
func Goodbye() {
    fmt.Println("Goodbye!")
}
```

**Organizzazione della directory**
```C++
myproject/
│
├── greeting/
│   ├── greet.go
│   └── farewell.go
└── main.go
```

**File main.go**
Il file main.go utilizzerà entrambe le funzioni Hello() e Goodbye() del package greeting:
```go
// main.go
package main

import (
    "greeting"
)

func main() {
    greeting.Hello()
    greeting.Goodbye()
}
```

## Compilazione e esecuzione
Poiché Go gestisce automaticamente l'importazione di tutti i file all'interno di un package, non devi preoccuparti di importare manualmente ogni singolo file. Quando esegui il comando per compilare il tuo progetto:
```BASH
go run main.go
```
Go compilerà tutti i file nel package greeting e li utilizzerà nel tuo programma.