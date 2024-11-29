package worker

import (
	"backend/internal/model"
	"backend/internal/service"
	"log/slog"
	"time"
)

func StartComputerWorker() {
	slog.Info("Starting computer worker...")

	for range time.Tick(1 * time.Minute) {
		if !service.ComputerData.Online {
			service.AddComputerData(model.ComputerWebSocketMessage{
				Cpu: 0,
				Ram: 0,
			})
		}
	}
}
