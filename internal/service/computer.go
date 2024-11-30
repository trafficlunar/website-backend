package service

import (
	"backend/internal/model"
	"time"
)

var ComputerData model.ComputerData = model.ComputerData{
	Online: false,
	Graph:  initializeGraphData(),
}

func initializeGraphData() []model.ComputerGraphData {
	graphData := make([]model.ComputerGraphData, 50)

	for i := 0; i < 50; i++ {
		graphData[i] = model.ComputerGraphData{
			Timestamp: time.Now().Truncate(1 * time.Minute).Add(time.Duration(-50+i) * time.Minute),
			Cpu:       0,
			Ram:       0,
		}
	}

	return graphData
}

func AddComputerData(clientMessage model.ComputerWebSocketMessage) {
	ComputerData.Graph = append(ComputerData.Graph, model.ComputerGraphData{
		Timestamp: time.Now().Truncate(time.Minute).Add(-time.Minute),
		Cpu:       int(clientMessage.Cpu),
		Ram:       int(clientMessage.Ram),
	})

	if len(ComputerData.Graph) > 50 {
		ComputerData.Graph = ComputerData.Graph[1:]
	}
}
