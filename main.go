package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hola, Mundo!")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	port := "8080"
	fmt.Printf("Escuchando en el puerto %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
