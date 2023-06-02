package handler

import (
	"context"
	"encoding/json"
	"go-backend/interfaces"
	"go-backend/interfaces/queue"
	"go-backend/model"
	"net/http"
	"time"
)

type WithdrawalHandler struct {
	interfaces.WithdrawalService
	queue.Publisher
}

func (h *WithdrawalHandler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var withdrawRequest model.WithdrawalRequest

		if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := h.Publisher.Publish(ctx, "withdraw_request", nil) // TODO WIP
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
