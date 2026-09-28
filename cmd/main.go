package main

import (
	"fmt"
	"log"
	"net/http"

	server "github.com/Robot-tim1/inspiring-quote-gen"
)

func main() {
	port := ":8080"
	fmt.Printf("Starting server on http://localhost%s/\n", port)

	httpServer := &http.Server{
		Handler: server.GetMux(),
		Addr:    port,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
