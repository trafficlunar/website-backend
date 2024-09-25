package worker

func StartWorkers() {
	go StartLastFMWorker()
}
