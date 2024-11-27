package worker

import "os"

func StartWorkers() {
	go StartLastFMWorker()
	if os.Getenv("STATUS") == "true" {
		go StartStatusWorker()
	}
}
