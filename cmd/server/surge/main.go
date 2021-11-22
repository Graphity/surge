package main

import (
	"fmt"
	"github.com/Graphity/surge/server/app"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")

	app.RegisterAllHandlers()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
