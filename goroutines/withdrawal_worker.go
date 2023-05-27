package goroutines

import (
	"fmt"
	"go-backend/model"
	"go-backend/service"
)

type Worker struct {
	Id                  int
	Work                chan model.WithdrawRequest
	WithdrawWorkerQueue chan chan model.WithdrawRequest
	QuitChan            chan bool
	withdrawalService   service.WithdrawalService
}

func NewWorker(id int, workerQueue chan chan model.WithdrawRequest, withdrawalService service.WithdrawalService) Worker {
	return Worker{
		Id:                  id,
		Work:                make(chan model.WithdrawRequest),
		WithdrawWorkerQueue: workerQueue,
		QuitChan:            make(chan bool),
		withdrawalService:   withdrawalService,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WithdrawWorkerQueue <- w.Work

			select {
			case withdrawRequest := <-w.Work:
				fmt.Println("Worker will start withdrawing")

				if err := w.withdrawalService.Withdraw(withdrawRequest); err != nil {
					fmt.Errorf("an error ocurred when withdrawing %g", err)
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
