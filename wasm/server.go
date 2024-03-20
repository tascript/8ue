package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./assets/")))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Failed to listen HTTP Server: %v", err)
	}
}
