package main

import (
	"log"
	"net/http"

	"github.com/israeloluwasegun293-stars/L2E-Learning/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)

	log.Println("server starting at http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
