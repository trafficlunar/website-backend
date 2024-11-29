package model

type ComputerWebsocketMessage struct {
	Cpu uint8 `json:"cpu"`
	Ram uint8 `json:"ram"`
}
