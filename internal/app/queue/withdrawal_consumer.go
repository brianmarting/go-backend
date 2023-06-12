package queue

import (
	"context"
	"encoding/json"
	"go-backend/internal/api/model"
	"go-backend/internal/observability/tracing"
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
			handleMessage(c, message)
		}
	}()
}

func handleMessage(c withdrawalConsumer, message Message) {
	tracer := tracing.GetTracer()
	_, span := tracer.Start(context.Background(), "receive-withdrawal-msg-queue")
	defer span.End()

	wr := &model.WithdrawalRequest{}

	if err := json.Unmarshal(message.GetBytes(), wr); err != nil {
		log.Info().Err(err).Msg("failed to parse data")
		if err = message.Nack(); err != nil {
			log.Info().Err(err).Msg("failed to nack msg")
		}
		return
	}

	err := c.withdrawalService.Withdraw(wr)
	if err != nil {
		log.Info().Err(err).Msg("failed to process withdraw request")
		if err = message.Nack(); err != nil {
			log.Info().Err(err).Msg("failed to nack msg")
		}
		return
	}

	if err = message.Ack(); err != nil {
		log.Info().Err(err).Msg("failed to ack")
	}
}
