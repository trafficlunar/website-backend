package handler

import (
	"backend/internal/service"
	"encoding/json"
	"net/http"
)

func HandleGetVisitorCounter(w http.ResponseWriter, r *http.Request) {
	data := service.GetVisitorCounter()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func HandlePatchVisitorCounter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}
