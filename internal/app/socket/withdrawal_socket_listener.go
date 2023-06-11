package socket

import (
	"context"
	"encoding/json"
	"go-backend/internal/api/model"
	"go-backend/internal/observability/tracing"
	"go-backend/internal/service"

	"github.com/rs/zerolog/log"
)

type WithdrawalSocketListener interface {
	Start()
}

type withdrawalSocketListener struct {
	listener          Listener
	withdrawalService service.WithdrawalService
}

func NewWithdrawalSocketListener(listener Listener, service service.WithdrawalService) WithdrawalSocketListener {
	return withdrawalSocketListener{
		listener:          listener,
		withdrawalService: service,
	}
}

func (w withdrawalSocketListener) Start() {
	done := make(chan interface{})

	in, out := w.listener.Start(done)

	go func() {
		defer func() {
			close(out)
			close(done)
		}()

		for msg := range in {
			handleMessage(msg, out, w)
		}
	}()
}

func handleMessage(msg Message, out chan<- string, w withdrawalSocketListener) {
	tracer := tracing.GetTracer()
	_, span := tracer.Start(context.Background(), "receive-withdrawal-msg-tcp-socket")
	defer span.End()

	wr := &model.WithdrawalRequest{}

	if err := json.Unmarshal(msg.GetBody(), wr); err != nil {
		log.Info().Err(err).Msg("failed to parse tcp inbound body, sending nack")
		out <- "NACK"
		return
	}

	err := w.withdrawalService.Withdraw(wr)
	if err != nil {
		log.Info().Err(err).Msg("failed to process withdraw request")
		out <- "NACK"
		return
	}

	out <- "ACK"
}
