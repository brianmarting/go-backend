package socket

import (
	"encoding/json"
	"go-backend/internal/api/model"
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
			wr := &model.WithdrawalRequest{}

			if err := json.Unmarshal(msg.GetBody(), wr); err != nil {
				log.Info().Err(err).Msg("failed to parse tcp inbound body")
				out <- "NACK"
				continue
			}

			err := w.withdrawalService.Withdraw(wr)
			if err != nil {
				log.Info().Err(err).Msg("failed to process withdraw request")
				out <- "NACK"
				continue
			}

			out <- "ACK"
		}
	}()
}
