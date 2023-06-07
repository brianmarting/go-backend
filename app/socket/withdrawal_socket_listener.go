package socket

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"go-backend/api/model"
	"go-backend/service"
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
	in, out, err := w.listener.Start()
	if err != nil {
		log.Info().Err(err).Msg("failed to start withdrawal listener")
		return
	}

	go func() {
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
