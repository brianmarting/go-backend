package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-backend/db"
	"go-backend/interfaces"
	"net/http"
)

type WalletHandler struct {
	Store interfaces.WalletStore
}

func (h *WalletHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		wallet, err := h.Store.Wallet(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(wallet)
	}
}

func (h *WalletHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wallet := &db.Wallet{
			Id:      uuid.New(),
			Address: uuid.NewString(),
		}

		if err := h.Store.CreateWallet(wallet); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
