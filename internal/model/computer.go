package model

import "time"

type ComputerWebSocketMessage struct {
	Cpu uint8 `json:"cpu"`
	Ram uint8 `json:"ram"`
}

type ComputerData struct {
	Online bool                `json:"online"`
	Graph  []ComputerGraphData `json:"graph"`
}

type ComputerGraphData struct {
	Timestamp time.Time `json:"timestamp"`
	Cpu       int       `json:"cpu"`
	Ram       int       `json:"ram"`
}
