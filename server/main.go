package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatalf("Failed to listen HTTP Server: %v", err)
	}
}
