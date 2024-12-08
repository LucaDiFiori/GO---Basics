In Go, un **modulo** è una collezione di package che vengono gestiti come un'unità singola. Serve per organizzare il codice e gestire le dipendenze di un progetto (i package esterni che utilizzi). Il sistema di moduli (`go mod`) ti aiuta a:

1. Specificare quali package esterni usi nel tuo progetto.
2. Scaricare e gestire automaticamente le versioni di quei package.

Ogni modulo è identificato da un **nome del modulo**, solitamente corrispondente a un URL o un percorso importabile, come `github.com/username/repository`.

***
## File principale: `go.mod`

Quando crei un modulo, Go genera un file chiamato `go.mod`. Questo file:

- Specifica il **nome del modulo** o il **percorso dove puà essere trovato**.
- Elenca le **dipendenze** (cioè altri moduli) necessari per il progetto.
- Gestisce le versioni delle dipendenze per garantire coerenza.

Esempio
```go
module github.com/luca/mio-progetto

go 1.20

require (
    github.com/gin-gonic/gin v1.9.0
    github.com/stretchr/testify v1.8.0
)
```



***
### Creare un modulo

Per creare un modulo in un nuovo progetto:

1. Apri il terminale nella directory del tuo progetto.
2. Esegui il comando:

```go
go mod init <nome-del-modulo>
```

Esempio:
```go
go mod init github.com/luca/mio-progetto
```

Questo creerà il file `go.mod` con il nome del modulo.



***
### Funzionalità principali dei moduli

1. **Gestione delle dipendenze**: Le dipendenze vengono specificate nel file `go.mod` e scaricate automaticamente tramite `go get`.
2. **Versionamento**: I moduli seguono le convenzioni semantiche di versioning (`v1.0.0`, `v2.1.3`, ecc.).
3. **Compatibilità**: Go garantisce che il modulo funzionerà con le versioni dichiarate nel file `go.mod`.
4. **Caching**: Le dipendenze scaricate vengono memorizzate nella cache del sistema, quindi non è necessario scaricarle di nuovo.



***
### File `go.sum`

Quando installi o aggiorni le dipendenze, Go genera un file `go.sum` accanto al file `go.mod`. Questo file:

- Elenca gli hash delle versioni delle dipendenze per verificare l'integrità.
- Aiuta a garantire che tutti gli sviluppatori abbiano le stesse dipendenze nel progetto.



***
### Struttura di un modulo

Ecco come potrebbe apparire un progetto Go con un modulo:

```go
mio-progetto/
├── go.mod
├── go.sum
├── main.go
└── util/
    ├── util.go
    └── util_test.go
```

- **`go.mod`**: Definisce il modulo e le sue dipendenze.
- **`go.sum`**: Tiene traccia delle versioni esatte delle dipendenze.
- **`main.go`**: File principale del programma.
- **`util/`**: Una cartella con pacchetti aggiuntivi