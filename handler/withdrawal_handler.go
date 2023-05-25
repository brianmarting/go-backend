package handler

import (
	"github.com/google/uuid"
	"go-backend/goroutines"
	"go-backend/interfaces"
	"go-backend/model"
	"net/http"
	"strconv"
)

type WithdrawalHandler struct {
	WalletStore       interfaces.WalletStore
	WalletCryptoStore interfaces.WalletCryptoStore
}

func (h *WithdrawalHandler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cryptoIdString := r.FormValue("cryptoId")

		cryptoId, err := uuid.Parse(cryptoIdString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fromAddress := r.FormValue("fromAddress")
		toAddress := r.FormValue("toAddress")
		amountString := r.FormValue("amount")
		amount, err := strconv.Atoi(amountString)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		withdrawRequest := model.WithdrawRequest{
			CryptoId:    cryptoId,
			FromAddress: fromAddress,
			ToAddress:   toAddress,
			Amount:      amount,
		}

		goroutines.WithdrawalQueue <- withdrawRequest
	}
}
