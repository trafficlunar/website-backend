package worker

func StartWorkers() {
	go StartLastFMWorker()
	go StartStatusWorker()
}
