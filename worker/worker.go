package worker

import (
	"github.com/rs/zerolog/log"
	"go-backend/interfaces"
)

type Worker struct {
	Id                      int
	WorkRequestChannel      chan interfaces.WorkerJob
	WorkRequestChannelQueue chan chan interfaces.WorkerJob
	QuitChan                chan bool
}

func NewWorker(id int, workRequestChannelQueue chan chan interfaces.WorkerJob, quit chan bool) Worker {
	return Worker{
		Id:                      id,
		WorkRequestChannel:      make(chan interfaces.WorkerJob),
		WorkRequestChannelQueue: workRequestChannelQueue,
		QuitChan:                quit,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkRequestChannelQueue <- w.WorkRequestChannel

			select {
			case workRequest := <-w.WorkRequestChannel:
				log.Info().Msg("will start working")

				if err := workRequest.Work(); err != nil {
					log.Err(err).Msg("an error occurred while working")
				}
			case <-w.QuitChan:
				log.Info().Msg("stopping worker")
				return
			}
		}
	}()
}
