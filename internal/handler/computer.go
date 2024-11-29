package handler

import (
	"backend/internal/model"
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

func HandleComputerWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error when upgrading websocket connection", slog.Any("error", err))
		return
	}
	defer conn.Close()

	slog.Info("Websocket connection established!")

	// Read messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			slog.Error("WebSocket connection closed by client", slog.Any("error", err))
			break
		}

		var clientMessage model.ComputerWebsocketMessage
		if err := json.Unmarshal(message, &clientMessage); err != nil {
			slog.Error("Error unmarshaling JSON", slog.Any("error", err))
			continue
		}

		slog.Info("Recieved message", slog.Any("message", clientMessage))
	}
}
