package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"go-backend/interfaces/db"
	"go-backend/interfaces/queue"
	"go-backend/model"
	"net/http"
	"time"
)

type WithdrawalHandler interface {
	Withdraw() http.HandlerFunc
}

type withdrawalHandler struct {
	publisher   queue.Publisher
	cryptoStore db.CryptoStore
	walletStore db.WalletStore
}

func NewWithdrawalHandler(publisher queue.Publisher, cryptoStore db.CryptoStore, walletStore db.WalletStore) WithdrawalHandler {
	return withdrawalHandler{
		publisher:   publisher,
		cryptoStore: cryptoStore,
		walletStore: walletStore,
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
	crypto, err := h.cryptoStore.GetByUuid(wr.CryptoId)
	if err != nil {
		return err
	}

	walletFrom, err := h.walletStore.GetByAddress(wr.FromAddress)
	if err != nil {
		return err
	}
	if crypto.Id != walletFrom.CryptoId {
		return errors.New("the wallet to address does not support the given cryptocurrency")
	}
	if walletFrom.Amount < wr.Amount {
		return errors.New("the wallet from does not have sufficient funds")
	}

	walletTo, err := h.walletStore.GetByAddress(wr.ToAddress)
	if err != nil {
		return err
	}
	if crypto.Id != walletTo.CryptoId {
		return errors.New("the wallet to address does not support the given cryptocurrency")
	}

	return nil
}
