package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
)

func GetMux() *http.ServeMux {
	mux := http.NewServeMux()
	registerRoutes(mux)
	return mux
}

func registerRoutes(mux *http.ServeMux) {
	fileServer := http.FileServer(http.Dir("./assets"))
	mux.Handle("GET /assets/", http.StripPrefix("/assets/", fileServer))

	mux.HandleFunc("GET /", handlerQuotes)
	mux.HandleFunc("GET /styles.css", handlerCssStyle)
}

type Quote struct {
	Quote    string `json:"quote"`
	Name     string `json:"name"`
	Filepath string `json:"filepath"`
}

func handlerQuotes(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("quotes.json")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	var quotes []Quote
	if err := json.NewDecoder(file).Decode(&quotes); err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	randomQuote := quotes[rand.N(len(quotes))]
	resp := fmt.Sprintf(`<body style="background-image: url('%s');"><h3 class="quote">%s<p class="quotesayer">- %s</p></h3></body>`, randomQuote.Filepath, randomQuote.Quote, randomQuote.Name)

	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	w.Header().Set("Link", `</styles.css>; rel="stylesheet"`)
	w.WriteHeader(200)
	w.Write([]byte(resp))
}

func handlerCssStyle(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("style.css")
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	css, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Error reading css stylesheet: %v", err)
	}

	w.Header().Set("Content-Type", "text/css")
	w.WriteHeader(200)
	w.Write(css)
}
