package worker

import (
	"backend/internal/model"
	"backend/internal/service"
	"log/slog"
	"time"
)

var StatusData model.StatusData

func StartStatusWorker() {
	slog.Info("Starting status worker...")
	StatusData = service.GetStatuses()

	for range time.Tick(5 * time.Minute) {
		slog.Info("Requesting Uptime Kuma...")
		StatusData = service.GetStatuses()
	}
}
