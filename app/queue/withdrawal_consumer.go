package queue

import (
	"encoding/json"
	"go-backend/api/model"
	"go-backend/service"

	"github.com/rs/zerolog/log"
)

type WithdrawalConsumer interface {
	StartConsuming()
}

type withdrawalConsumer struct {
	consumer          Consumer
	withdrawalService service.WithdrawalService
}

func NewWithdrawalConsumer(consumer Consumer, withdrawalService service.WithdrawalService) WithdrawalConsumer {
	return withdrawalConsumer{
		consumer:          consumer,
		withdrawalService: withdrawalService,
	}
}

func (c withdrawalConsumer) StartConsuming() {
	messages, err := c.consumer.StartConsuming("withdrawRequests", "withdraw.request")
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

			err := c.withdrawalService.Withdraw(wr)
			if err != nil {
				log.Info().Err(err).Msg("failed to process withdraw request")
			}

			message.Ack()
		}
	}()
}
