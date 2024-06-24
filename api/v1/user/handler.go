package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"twitter-clone/internal/api"
)

func GetUserById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Set("Content-Type", "application/json")

	if _, err := uuid.Parse(id); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode(api.ErrorResponse{Error: api.InnerError{
			Message: "Invalid UUID",
		}})
		if err != nil {
			log.Printf("Failed to encode response: %v\n", err)
		}
		return
	}

	if v, ok := mockUsers[id]; ok {
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(v)
		if err != nil {
			log.Printf("Failed to encode response: %v\n", err)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(api.ErrorResponse{Error: api.InnerError{Message: fmt.Sprintf("Unable to find user ID: %v", id)}})
	}
}
