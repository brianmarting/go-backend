package goroutines

import (
	"fmt"
	"go-backend/model"
	"go-backend/service"
)

var WithdrawalQueue = make(chan model.WithdrawRequest, 100)

var WithdrawalWorkerQueue chan chan model.WithdrawRequest

func StartDispatcher(amount int, withdrawalService service.WithdrawalService) {
	WithdrawalWorkerQueue = make(chan chan model.WithdrawRequest, 5)

	for i := 0; i < amount; i++ {
		worker := NewWorker(i, WithdrawalWorkerQueue, withdrawalService)
		worker.Start()
	}

	go func() {
		for {
			select {
			case withdrawal := <-WithdrawalQueue:
				fmt.Println("Incoming withdrawal req")
				// Start the withdrawal work
				go func() {
					// get idle worker from queue
					withdrawalWorker := <-WithdrawalWorkerQueue

					fmt.Println("Got idle worker from queue")

					// add withdrawal work to worker so it can process it
					withdrawalWorker <- withdrawal
				}()
			}
		}
	}()
}
