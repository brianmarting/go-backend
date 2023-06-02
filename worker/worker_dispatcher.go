package worker

import (
	"github.com/rs/zerolog/log"
	"go-backend/interfaces"
)

var WorkRequestChannel = make(chan interfaces.WorkerJob, 100)

var QuitWorkRequestChannel = make(chan bool)

var WorkRequestChannelQueue chan chan interfaces.WorkerJob

func StartDispatcher(amount int) {
	WorkRequestChannelQueue = make(chan chan interfaces.WorkerJob, 5)

	for i := 0; i < amount; i++ {
		worker := NewWorker(i, WorkRequestChannelQueue, QuitWorkRequestChannel)
		worker.Start()
	}

	go func() {
		defer func() {
			close(WorkRequestChannel)
			close(QuitWorkRequestChannel)
			close(WorkRequestChannelQueue)
		}()

		for {
			select {
			case workerJob := <-WorkRequestChannel:
				log.Info().Msg("Incoming withdrawal req")
				// Start the withdrawal work
				go func() {
					// get idle worker from queue
					workRequestChannel := <-WorkRequestChannelQueue

					log.Info().Msg("Got idle worker from queue")

					// add work to worker so it can process it
					workRequestChannel <- workerJob
				}()
			}
		}
	}()
}

func StopWorkers() {
	QuitWorkRequestChannel <- true
}
