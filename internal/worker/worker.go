package worker

import "os"

func StartWorkers() {
	go StartLastFMWorker()
	go StartComputerWorker()
	if os.Getenv("UPTIME_KUMA_ENABLED") == "true" {
		go StartStatusWorker()
	}
}
