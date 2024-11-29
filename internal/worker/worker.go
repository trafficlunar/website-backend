package worker

import "os"

func StartWorkers() {
	go StartLastFMWorker()
	if os.Getenv("UPTIME_KUMA_ENABLED") == "true" {
		go StartStatusWorker()
	}
}
