package main

import (
	"log"
	"net/http"

	"github.com/masterhorn/XM-Golang-Exercise-v21.0.0/internal"
)

func main() {
	r := internal.Handlers()

	log.Println("Starting server at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
