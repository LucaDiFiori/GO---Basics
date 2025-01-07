Le **slice** in Go sono una struttura fondamentale per gestire insiemi di dati in modo dinamico e flessibile. 
A differenza degli array, le slice **non hanno una dimensione fissa** e possono crescere o ridursi dinamicamente. 

Internamente, una slice non contiene direttamente i dati, ma è un'astrazione sopra un array esistente. Questo concetto ha alcune implicazioni fondamentali per il modo in cui le slice funzionano e vengono gestite:

### Struttura e caratteristiche di una Slice

- **Dinamiche:** A differenza degli array, le slice non hanno una dimensione fissa. Possono crescere o ridursi durante l'esecuzione del programma.
  
- **Collegamento a un array sottostante:** Una slice è un'astrazione su un array. Modificare un elemento della slice modifica anche l'elemento corrispondente nell'array sottostante.
  
- **Campi principali:** Una slice è definita da:
    - **Puntatore:** Indica l'inizio dell'array sottostante.
    - **Lunghezza:** Numero di elementi attualmente nella slice.
    - **Capacità:** Numero massimo di elementi che la slice può contenere senza riallocare memoria.



***
## Dichiarazione di una slice

#### 1. Creare una slice vuota
```go
var s []int // Una slice di interi inizialmente vuota
```

#### 2. Inizializzare letterale di una slice
```go
s := []int{1, 2, 3, 4}
```



***
## Lunghezza e capacità
- La **lunghezza** (`len`) della slice è il numero di elementi nella finestra.
- La **capacità** (`cap`) è calcolata dal punto di inizio della slice fino alla fine dell'array sottostante.

```go
arr := [5]int{10, 20, 30, 40, 50}

slice := arr[1:4]
fmt.Println(len(slice)) // Lunghezza: 3 (20, 30, 40)
fmt.Println(cap(slice)) // Capacità: 4 (20, 30, 40, 50)
```




***
## Creazione una slice da un array esistente
Quando crei una slice da un array esistente in Go, stai definendo una "finestra" che fa riferimento a un sottoinsieme degli elementi dell'array originale. La slice **non copia** i dati; invece, condivide la stessa memoria con l'array sottostante. Questo rende le slice efficienti ma anche potenzialmente pericolose se modifichi i dati.

#### Sintassi
```go
slice := array[start:end]
```

Esempio
```go
arr := [5]int{10, 20, 30, 40, 50}

slice := arr[1:4] // Include gli elementi da indice 1 a 3
fmt.Println(slice) // Output: [20 30 40]
```

- **Indice iniziale:** `1` corrisponde al secondo elemento (`20`).
- **Indice finale:** `4` è esclusivo, quindi l'elemento con indice `4` (`50`) **non** è incluso.

Quello che stiamo cioè facendo è **prendere tutti gli elementi dal primo al quarto escluso**


#### Caratteristiche

**Condivisione della memoria:**
La slice punta all'array sottostante. Modificando gli elementi della slice, modifichi anche l'array originale.

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]
slice[0] = 25 // Modifica l'array originale
fmt.Println(arr) // Output: [10 25 30 40 50]

```


**Slice di una slice**
Puoi creare una nuova slice a partire da una slice esistente.

```go
arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]    // [20, 30, 40]
subSlice := slice[1:] // [30, 40]
fmt.Println(subSlice) // Output: [30 40]
```

**NOTA**: Quando in Go scrivi `slice[start:]` (cioè con i due punti `:` ma senza specificare un indice finale), stai creando una nuova slice che include tutti gli elementi dalla posizione `start` fino alla **fine della slice originale**.



***
***
## Creazione una slice - make
In Go, puoi creare una slice utilizzando la funzione built-in **`make`**, che ti permette di inizializzare una slice con una dimensione e, opzionalmente, una capacità specifica. Questa funzione è utile quando hai bisogno di una slice di una certa lunghezza senza doverla popolare manualmente o quando vuoi ottimizzare la memoria.

#### Sintassi
```go
slice := make([]T, len, cap)
```

- **`T`**: Tipo degli elementi della slice (es. `int`, `string`, ecc.).
- **`len`**: Lunghezza iniziale della slice (numero di elementi già allocati e inizializzati).
- **`cap`** (opzionale): Capacità totale della slice. Se omesso, `cap` sarà uguale a `len`.

#### Quando usare `make` per le slice?
- **Quando hai bisogno di una slice vuota ma con una lunghezza o capacità specifica.**
- **Quando vuoi ottimizzare le prestazioni, preallocando memoria per evitare riallocazioni frequenti.**


#### Esempi
**Creare una slice con lunghezza specifica**
```go
s := make([]int, 5) // Slice di 5 elementi, capacità 5
fmt.Println(s)      // Output: [0 0 0 0 0]
```
- La slice `s` contiene 5 elementi inizializzati al **valore zero** del tipo (`0` per `int`).



**Creare una slice con lunghezza e capacità diverse**
```go
s := make([]int, 3, 5) // Slice di lunghezza 3, capacità 5

fmt.Println(s)         // Output: [0 0 0]
fmt.Println(len(s))    // Output: 3
fmt.Println(cap(s))    // Output: 5
```
- La slice `s` ha 3 elementi utilizzabili inizialmente.
- La capacità totale (`cap`) permette di aggiungere fino a 2 elementi senza riallocare memoria.




****

## Aggiungere elementi ad una slice

Puoi usare `append` per aggiungere elementi:

Esempio 1:
```go
s := make([]int, 3, 5) // Lunghezza 3, capacità 5
s = append(s, 10, 20)  // Aggiunge 2 elementi

fmt.Println(s)         // Output: [0 0 0 10 20]
```

- - **Lunghezza (`len`)**: 3 → Contiene 3 elementi inizializzati a `0`.
    - **Capacità (`cap`)**: 5 → Ha spazio per un massimo di 5 elementi senza riallocare memoria.
- Valore iniziale della slice: `[0, 0, 0, _, _]`.

- `append` aggiunge gli elementi `10` e `20` alla slice: `[0, 0, 0, 10, 20]`.

- Dopo l'operazione:
	- La **lunghezza** diventa `5` (3 elementi iniziali + 2 aggiunti).
	- La **capacità** rimane `5`, perché non è stata superata.


Se aggiungi più elementi di quanti la capacità permetta, Go rialloca automaticamente la memoria.
```go
s := make([]int, 2, 4) // Lunghezza 2, capacità 4
s = append(s, 10, 20, 30) // Aggiunge 3 elementi, supera la capacità
fmt.Println(s)            // Output: [0 0 10 20 30]
fmt.Println(cap(s))       // Capacità maggiore della precedente
```




Esempio 2:
```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)
fmt.Println(slice) // Output: [1 2 3 4 5]
```




***
### **Copiare una slice**

Puoi copiare i contenuti di una slice in un'altra usando la funzione built-in `copy`:
```go
src := []int{10, 20, 30}
dest := make([]int, len(src))
copy(dest, src)
fmt.Println(dest) // Output: [10 20 30]
```





***
### Esempio
```go
package main

import "fmt"

func main() {
    // Creazione di una slice da un array
    arr := [5]int{10, 20, 30, 40, 50}
    slice := arr[1:4]

    fmt.Println("Slice:", slice)         // Output: [20 30 40]
    fmt.Println("Length:", len(slice))  // Output: 3
    fmt.Println("Capacity:", cap(slice)) // Output: 4

    // Aggiunta di elementi
    slice = append(slice, 60)
    fmt.Println("Updated Slice:", slice) // Output: [20 30 40 60]

    // Iterazione
    for index, value := range slice {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    // Copia di una slice
    newSlice := make([]int, len(slice))
    copy(newSlice, slice)
    fmt.Println("Copied Slice:", newSlice) // Output: [20 30 40 60]
}
```

Il **`range`** in Go è un **costrutto che consente di iterare** su vari tipi di collezioni come slice, array, map, string e canali. Quando utilizzi `range` in un loop, fornisce una coppia di valori per ogni elemento della collezione: l'indice (o chiave) e il valore corrispondente.

### Cos'è `range`?
Il **`range`** in Go è un **costrutto che consente di iterare** su vari tipi di collezioni come slice, array, map, string e canali. Quando utilizzi `range` in un loop, fornisce una coppia di valori per ogni elemento della collezione: l'indice (o chiave) e il valore corrispondente.

`range` è una parola chiave che:

1. **Si "aggancia" a una collezione** (come una slice o un array).
2. **Restituisce automaticamente un valore alla volta**, insieme al relativo indice (o chiave).

In altre parole:

- Per una slice o un array, restituisce **l'indice** e **il valore** di ogni elemento.
- Per una mappa, restituisce **la chiave** e **il valore**.
- Per una stringa, restituisce **l'indice** e **il carattere (in formato rune)**.
- Per un canale, restituisce il **valore** ricevuto dal canale.



