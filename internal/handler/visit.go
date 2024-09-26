package handler

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

func HandleGetVisitCounter(w http.ResponseWriter, r *http.Request) {
	data := service.GetVisitCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandlePatchVisitCounter(w http.ResponseWriter, r *http.Request) {
	data := service.IncrementVisitCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
