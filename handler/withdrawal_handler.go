package handler

import (
	"encoding/json"
	"go-backend/goroutines"
	"go-backend/interfaces"
	"go-backend/model"
	"net/http"
)

type WithdrawalHandler struct {
	WalletStore       interfaces.WalletStore
	WalletCryptoStore interfaces.WalletCryptoStore
}

func (h *WithdrawalHandler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var withdrawRequest model.WithdrawalRequest

		if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		goroutines.WithdrawalRequestChannel <- withdrawRequest
	}
}
