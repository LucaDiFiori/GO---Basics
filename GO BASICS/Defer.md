`defer` è una struttura speciale per ritardare l'esecuzione di un'istruzione fino alla fine della funzione che la contiene.

**Esempio:**
```go
package main

import "fmt"

func main() {
    defer fmt.Println("This is deferred.")
    fmt.Println("This is printed first.")
}
```

Output:
```go
This is printed first.
This is deferred.
```

I comandi `defer` vengono eseguiti in ordine **LIFO** (Last In, First Out).