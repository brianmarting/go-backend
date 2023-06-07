package handlers

import (
	"encoding/json"
	"go-backend/interfaces/db"
	"go-backend/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type WalletHandler interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
}

type walletHandler struct {
	store db.WalletStore
}

func NewWalletHandler(store db.WalletStore) WalletHandler {
	return walletHandler{
		store: store,
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

		wallet, err := h.store.GetByUuid(id)
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

		if err := h.store.Create(wallet); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
