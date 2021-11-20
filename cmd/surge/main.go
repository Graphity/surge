package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	indexPath := "./web/static/index.html"
	tmpl := template.Must(template.ParseFiles(indexPath))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong\n")
	if err != nil {
		return
	}
}

func main() {
	fmt.Println("Starting server...")

	http.HandleFunc("/", index)
	http.HandleFunc("/ping", ping)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
