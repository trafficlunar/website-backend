package handler

import (
	"backend/internal/model"
	"backend/internal/service"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// todo: change for security
		return true
	},
}

func HandleComputerWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error when upgrading websocket connection", slog.Any("error", err))
		return
	}
	defer conn.Close()

	slog.Info("Websocket connection established!")
	online := true

	// Read messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("WebSocket connection closed by client", slog.Any("error", err))
			online = false
			break
		}

		var clientMessage model.ComputerWebSocketMessage
		if err := json.Unmarshal(message, &clientMessage); err != nil {
			slog.Error("Error unmarshalling JSON", slog.Any("error", err))
			continue
		}

		service.AddComputerData(online, clientMessage)
		slog.Info("Recieved message", slog.Any("message", clientMessage))
	}
}

func HandleComputerGraphData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(service.ComputerData)
}
