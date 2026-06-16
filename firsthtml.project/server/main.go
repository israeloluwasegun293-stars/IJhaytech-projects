package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	log.Println("server starting at http://localhost/8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
