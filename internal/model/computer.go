package model

type ComputerWebSocketMessage struct {
	Cpu uint8 `json:"cpu"`
	Ram uint8 `json:"ram"`
}

type ComputerGraphData struct {
	Online bool  `json:"online"`
	Cpu    []int `json:"cpu"`
	Ram    []int `json:"ram"`
}
