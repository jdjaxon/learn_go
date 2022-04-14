package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const Port = "8080"
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server listening on port %s\n", Port)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
} // main

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello!")
} // helloHandler

func formHandler(w http.ResponseWriter, r *http.Request) {
} // formHandler
