package goroutines

import (
	"github.com/goioc/di"
	"github.com/rs/zerolog/log"
	"go-backend/model"
	"go-backend/service"
)

type Worker struct {
	Id                  int
	Work                chan model.WithdrawRequest
	WithdrawWorkerQueue chan chan model.WithdrawRequest
	QuitChan            chan bool
	withdrawalService   *service.WithdrawalService
}

func NewWorker(id int, workerQueue chan chan model.WithdrawRequest) Worker {
	return Worker{
		Id:                  id,
		Work:                make(chan model.WithdrawRequest),
		WithdrawWorkerQueue: workerQueue,
		QuitChan:            make(chan bool),
		withdrawalService:   di.GetInstance("withdrawalService").(*service.WithdrawalService),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WithdrawWorkerQueue <- w.Work

			select {
			case withdrawRequest := <-w.Work:
				log.Info().Msg("Worker will start withdrawing")

				if err := w.withdrawalService.Withdraw(withdrawRequest); err != nil {
					log.Err(err).Msg("an error occurred when withdrawing")
				}
			case <-w.QuitChan:
				return
			}
		}
	}()
}

// Stop the work non-blocking
func (w *Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
