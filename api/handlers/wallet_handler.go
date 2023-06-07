package handlers

import (
	"encoding/json"
	"go-backend/persistence/db/model"
	"go-backend/service"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type WalletHandler interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
}

type walletHandler struct {
	service service.WalletService
}

func NewWalletHandler(service service.WalletService) WalletHandler {
	return walletHandler{
		service: service,
	}
}

func (h walletHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wallet, err := h.service.GetByUuid(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wallet)
	}
}

func (h walletHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet := model.Wallet{
			Uuid:    uuid.New(),
			Address: uuid.NewString(),
		}

		if err := h.service.Create(wallet); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
