package handler

import (
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
	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Error("Error when upgrading websocket connection", slog.Any("error", err))
		return
	}
	defer connection.Close()

	slog.Info("Websocket connection established!")

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			slog.Error("Error reading WebSocket message", slog.Any("error", err))
			break
		}
		slog.Info("Recieved message", slog.Any("message", message))
	}
}
