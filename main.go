package main

import (
	"fmt"
	"log"
	"net/http"
)

func demo(w http.ResponseWriter, r *http.Request) {
	path := "./static/index.html"
	http.ServeFile(w, r, path)
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", demo)
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}
}
