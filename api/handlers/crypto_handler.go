package handlers

import (
	"encoding/json"
	"go-backend/interfaces/db"
	"go-backend/model"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type CryptoHandler interface {
	Get() http.HandlerFunc
	Create() http.HandlerFunc
	Delete() http.HandlerFunc
}

type cryptoHandler struct {
	store db.CryptoStore
}

func NewCryptoHandler(store db.CryptoStore) CryptoHandler {
	return cryptoHandler{
		store: store,
	}
}

func (c cryptoHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		crypto, err := c.store.GetByUuid(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crypto)
	}
}

func (c cryptoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		crypto := model.Crypto{
			Uuid: uuid.New(),
			Name: r.FormValue("name"),
		}

		if err := c.store.Create(crypto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}

func (c cryptoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.store.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
