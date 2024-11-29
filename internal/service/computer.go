package service

import (
	"backend/internal/model"
)

var ComputerData model.ComputerGraphData = model.ComputerGraphData{
	Cpu: make([]int, 20),
	Ram: make([]int, 20),
}

func AddComputerData(clientMessage model.ComputerWebSocketMessage) {
	ComputerData.Cpu = append(ComputerData.Cpu, int(clientMessage.Cpu))
	ComputerData.Ram = append(ComputerData.Ram, int(clientMessage.Ram))

	if len(ComputerData.Cpu) > 20 {
		ComputerData.Cpu = ComputerData.Cpu[1:]
		ComputerData.Ram = ComputerData.Ram[1:]
	}
}
