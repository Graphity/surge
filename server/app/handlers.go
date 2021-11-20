package app

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	indexPath := "./web/static/index.html"
	tmpl := template.Must(template.ParseFiles(indexPath))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func Ping(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "pong\n")
	if err != nil {
		return
	}
}