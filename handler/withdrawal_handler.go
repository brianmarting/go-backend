package handler

import (
	"encoding/json"
	"go-backend/interfaces"
	"go-backend/model"
	"go-backend/worker"
	"net/http"
)

type WithdrawalHandler struct {
	interfaces.WithdrawalService `di.inject:"withdrawalService"`
}

func (h *WithdrawalHandler) Withdraw() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var withdrawRequest model.WithdrawalRequest

		if err := json.NewDecoder(r.Body).Decode(&withdrawRequest); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		job := &worker.WithdrawalWorkerJob{
			WorkFn: func() error {
				return h.WithdrawalService.Withdraw(withdrawRequest)
			},
		}
		worker.WorkRequestChannel <- job
	}
}
