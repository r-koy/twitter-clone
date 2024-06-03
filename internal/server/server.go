package server

import (
	"log"
	"net/http"
)

func Start() {
	r := NewRouter()

	log.Println("Starting server on :3000")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
