package queue

import (
	"encoding/json"
	"fmt"
	"go-backend/facade"
	"go-backend/interfaces/queue"
	"go-backend/interfaces/service"
	"go-backend/model"
	serviceImpl "go-backend/service"
	"os"
	"sync"

	"github.com/rs/zerolog/log"
)

var once sync.Once

var consumer queue.Consumer

type WithdrawalConsumer struct {
	queue.Consumer
	service.WithdrawalService
}

func init() {
	once.Do(func() {
		var (
			username = os.Getenv("RABBITMQ_USERNAME")
			password = os.Getenv("RABBITMQ_PASSWORD")
			host     = os.Getenv("RABBITMQ_HOST")
		)
		consumer = facade.NewConsumer(fmt.Sprintf("amqp://%s:%s@%s:5672/", username, password, host))

		withdrawalConsumer := WithdrawalConsumer{
			Consumer:          consumer,
			WithdrawalService: serviceImpl.GetWithdrawalService(),
		}
		startConsuming(withdrawalConsumer)
	})
}

func startConsuming(withdrawalConsumer WithdrawalConsumer) {
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
}
