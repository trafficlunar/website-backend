package handler

import (
	"backend/internal/worker"
	"encoding/json"
	"net/http"
)

func HandleGetStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worker.StatusData)
}
