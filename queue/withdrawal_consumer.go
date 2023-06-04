package queue

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"go-backend/facade"
	"go-backend/interfaces/queue"
	"go-backend/interfaces/service"
	"go-backend/model"
	serviceImpl "go-backend/service"
	"sync"
)

var once sync.Once

var consumer queue.Consumer

type WithdrawalConsumer struct {
	queue.Consumer
	service.WithdrawalService
}

func init() {
	once.Do(func() {
		consumer = facade.NewConsumer("amqp://guest:guest@localhost:5672/")

		withdrawalConsumer := WithdrawalConsumer{
			Consumer:          consumer,
			WithdrawalService: serviceImpl.GetWithdrawalService(),
		}

		messages, err := withdrawalConsumer.StartConsuming("withdrawRequests", "withdraw.request")
		if err != nil {
			log.Fatal().Msg("failed to start consuming")
		}

		go func() {
			for message := range messages {
				wr := &model.WithdrawalRequest{}

				if err := json.Unmarshal(message.GetBytes(), wr); err != nil {
					log.Info().Err(err).Msg("failed to parse data")
					return
				}

				err := withdrawalConsumer.Withdraw(wr)
				if err != nil {
					log.Info().Err(err).Msg("failed to process withdraw request")
				}

				message.Ack()
			}
		}()
	})
}
