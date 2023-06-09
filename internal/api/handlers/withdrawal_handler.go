package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"go-backend/internal/api/model"
	"go-backend/internal/app/queue"
	"go-backend/internal/service"
	"net/http"
	"time"
)

type WithdrawalHandler interface {
	Withdraw() http.HandlerFunc
}

type withdrawalHandler struct {
	publisher     queue.Publisher
	cryptoService service.CryptoService
	walletService service.WalletService
}

func NewWithdrawalHandler(
	publisher queue.Publisher,
	cryptoService service.CryptoService,
	walletService service.WalletService,
) WithdrawalHandler {
	return withdrawalHandler{
		publisher:     publisher,
		cryptoService: cryptoService,
		walletService: walletService,
	}
}

func (h withdrawalHandler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var withdrawRequest model.WithdrawalRequest

		if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := validateWithdrawRequest(h, withdrawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wrBytes, err := json.Marshal(withdrawRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err = h.publisher.Publish(ctx, "withdraw.request", wrBytes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}

func validateWithdrawRequest(h withdrawalHandler, wr model.WithdrawalRequest) error {
	crypto, err := h.cryptoService.GetByUuid(wr.CryptoId)
	if err != nil {
		return err
	}

	walletFrom, err := h.walletService.GetByAddress(wr.FromAddress)
	if err != nil {
		return err
	}
	if crypto.Id != walletFrom.CryptoId {
		return errors.New("the wallet to address does not support the given cryptocurrency")
	}
	if walletFrom.Amount < wr.Amount {
		return errors.New("the wallet from does not have sufficient funds")
	}

	walletTo, err := h.walletService.GetByAddress(wr.ToAddress)
	if err != nil {
		return err
	}
	if crypto.Id != walletTo.CryptoId {
		return errors.New("the wallet to address does not support the given cryptocurrency")
	}

	return nil
}
