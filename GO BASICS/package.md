In Go, un **package** è un insieme di codice organizzato in modo logico e riutilizzabile. I pacchetti sono fondamentali per strutturare, gestire e mantenere il codice in un progetto Go. Ecco i concetti principali:

***
### **Cosa sono i package?**

Un package è un insieme di file sorgente `.go` che condividono lo stesso spazio dei nomi. Ogni file Go appartiene a un package, dichiarato nella parte superiore del file con la parola chiave `package`:

```go
package nomeDelPackage
```

ad esempio:
```go
package main
```


***
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




***
## Tipi i package

### 1. Package principali (eseguibili)
Il package speciale `main` è il punto di partenza di un'applicazione Go. Deve includere una funzione `main()` che viene eseguita quando il programma parte

esempio
```go
package main

import "fmt"

func main() {
    fmt.Println("Ciao, mondo!")
}
```


### 2. Package libreria
Contengono funzioni e strutture utili che possono essere importate e utilizzate da altri pacchetti.

Esempio: Il package standard `math`
```go
import "math"

func main() {
    fmt.Println(math.Sqrt(16)) // Calcola la radice quadrata
}
```



***
## **Importazione dei package**
Go fornisce un'ampia libreria standard con pacchetti predefiniti che coprono operazioni comuni come input/output, gestione dei file, calcoli matematici, ecc. 

Puoi importarli con la parola chiave `import`:
```go
import "nomeDelPackage"
```


È possibile importare più pacchetti in blocco:

```go
import (
    "fmt"
    "math"
)
```




***
## Organizzazione dei file

I package possono essere divisi in più file, ma tutti i file devono dichiarare lo stesso nome di package in cima. Ad esempio, tutti i file che fanno parte del package utils devono iniziare con:

```go

package utils

```


### Directory
In Go, il concetto di **package** è strettamente legato alla struttura delle directory nel file system. Ogni directory rappresenta un **package**, e tutti i file sorgente (`.go`) all'interno di quella directory devono dichiarare di appartenere allo stesso package.

**Esempio**:
Creiamo un programma che utilizza un **package personalizzato** chiamato `greeting`. Seguiremo questi passaggi:

1. **Creazione del package `greeting`**:  
    Nella directory `greeting`, creeremo due file, `greet.go` e `farewell.go`, che comporranno il package. Entrambi i file apparterranno al package `greeting`, dichiarato all'inizio di ciascun file.
    
2. **Utilizzo del package `greeting`**:  
    Scriveremo un programma principale (`main.go`) che importa e utilizza il package `greeting`.


Organizzazione della directory
```C++
myproject/
│
├── greeting/
│   ├── greet.go
│   └── farewell.go
└── main.go
```


File grert.go
```go
// greet.go
package greeting

import "fmt"

// Funzione esportata (pubblica)
func Hello() {
    fmt.Println("Hello!")
}
```


file farewell.go
```go
// farewell.go
package greeting

import "fmt"

// Funzione esportata (pubblica)
func Goodbye() {
    fmt.Println("Goodbye!")
}
```

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



### Compilazione e esecuzione
Poiché Go gestisce automaticamente l'importazione di tutti i file all'interno di un package, non devi preoccuparti di importare manualmente ogni singolo file. Quando esegui il comando per compilare il tuo progetto:

```BASH
go run main.go
```

Go compilerà tutti i file nel package greeting e li utilizzerà nel tuo programma.





***
## Creare un package personalizzato
Puoi creare i tuoi package per organizzare il codice.  
Supponiamo di avere una struttura del progetto così:

```iua
mio_progetto/
|-- main.go
|-- calcoli/
    |-- calcoli.go
```


1. **Definisci il package** nel file `calcoli/calcoli.go`:

```go
package calcoli

func Somma(a, b int) int {
    return a + b
}
```


2. **Usa il package in `main.go`**:
```go
package main

import (
    "fmt"
    "mio_progetto/calcoli" // Importa il package personalizzato
)

func main() {
    risultato := calcoli.Somma(3, 4)
    fmt.Println("Risultato:", risultato)
}
```





***
## Visibilità (esportazione)
In Go, le funzioni, variabili e tipi definiti in un package sono visibili all'esterno solo se il loro nome inizia con una **lettera maiuscola**.

- **Pubblico (esportato):** `Somma`, `MathFunction`
- **Privato:** `sommaPrivata`, `_helper`


Esempio:
```go
// Pubblico
func Somma(a, b int) int {
    return a + b
}

// Privato
func sommaPrivata(a, b int) int {
    return a + b
}
```




***
