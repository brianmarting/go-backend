package queue

import (
	"encoding/json"
	"go-backend/internal/api/model"
	"go-backend/internal/service"

	"github.com/rs/zerolog/log"
)

type WithdrawalConsumer interface {
	Start()
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

func (c withdrawalConsumer) Start() {
	messages, err := c.consumer.StartConsuming("withdrawRequests", "withdraw.request")
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start consuming")
		return
	}

	go func() {
		for message := range messages {
			wr := &model.WithdrawalRequest{}

			if err := json.Unmarshal(message.GetBytes(), wr); err != nil {
				log.Info().Err(err).Msg("failed to parse data")
				if err = message.Nack(); err != nil {
					log.Info().Err(err).Msg("failed to nack msg")
				}
				continue
			}

			err := c.withdrawalService.Withdraw(wr)
			if err != nil {
				log.Info().Err(err).Msg("failed to process withdraw request")
				if err = message.Nack(); err != nil {
					log.Info().Err(err).Msg("failed to nack msg")
				}
				continue
			}

			if err = message.Ack(); err != nil {
				log.Info().Err(err).Msg("failed to ack")
			}
		}
	}()
}
