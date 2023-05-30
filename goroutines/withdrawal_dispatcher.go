package goroutines

import (
	"github.com/rs/zerolog/log"
	"go-backend/model"
)

var WithdrawalRequestChannel = make(chan model.WithdrawalRequest, 100)

var WithdrawalRequestChannelQueue chan chan model.WithdrawalRequest

func StartDispatcher(amount int) {
	WithdrawalRequestChannelQueue = make(chan chan model.WithdrawalRequest, 5)

	for i := 0; i < amount; i++ {
		worker := NewWorker(i, WithdrawalRequestChannelQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case withdrawal := <-WithdrawalRequestChannel:
				log.Info().Msg("Incoming withdrawal req")
				// Start the withdrawal work
				go func() {
					// get idle worker from queue
					withdrawalRequestChannel := <-WithdrawalRequestChannelQueue

					log.Info().Msg("Got idle worker from queue")

					// add withdrawal work to worker so it can process it
					withdrawalRequestChannel <- withdrawal
				}()
			}
		}
	}()
}
