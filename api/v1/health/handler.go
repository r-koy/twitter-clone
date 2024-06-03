package health

import (
	"encoding/json"
	"net/http"
)

func Ping(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]string{"ping": "pong"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
