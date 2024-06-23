package server

import (
	"log"
	"net/http"
)

func Start() {
	r := NewRouter()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
