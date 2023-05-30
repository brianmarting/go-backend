package goroutines

import (
	"github.com/goioc/di"
	"github.com/rs/zerolog/log"
	"go-backend/model"
	"go-backend/service"
)

type Worker struct {
	Id                            int
	WithdrawalRequestChannel      chan model.WithdrawalRequest
	WithdrawalRequestChannelQueue chan chan model.WithdrawalRequest
	QuitChan                      chan bool
	withdrawalService             *service.WithdrawalService
}

func NewWorker(id int, WithdrawalRequestChannelQueue chan chan model.WithdrawalRequest) Worker {
	return Worker{
		Id:                            id,
		WithdrawalRequestChannel:      make(chan model.WithdrawalRequest),
		WithdrawalRequestChannelQueue: WithdrawalRequestChannelQueue,
		QuitChan:                      make(chan bool),
		withdrawalService:             di.GetInstance("withdrawalService").(*service.WithdrawalService),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WithdrawalRequestChannelQueue <- w.WithdrawalRequestChannel

			select {
			case withdrawRequest := <-w.WithdrawalRequestChannel:
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
