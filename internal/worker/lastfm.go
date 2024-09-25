package worker

import (
	"backend/internal/model"
	"backend/internal/service"
	"log/slog"
	"time"
)

var LastFMData model.LastFMData

func StartLastFMWorker() {
	slog.Info("Starting last.fm worker...")
	LastFMData = service.GetLastFMData()

	for range time.Tick(30 * time.Second) {
		slog.Info("Requesting last.fm...")
		LastFMData = service.GetLastFMData()
	}
}
