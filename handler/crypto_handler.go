package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go-backend/db"
	"go-backend/interfaces"
	"net/http"
)

type CryptoHandler struct {
	Store interfaces.CryptoStore
}

func (c *CryptoHandler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		crypto, err := c.Store.Crypto(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(crypto)
	}
}

func (c *CryptoHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		crypto := &db.Crypto{
			Id:   uuid.New(),
			Name: r.FormValue("name"),
		}

		if err := c.Store.CreateCrypto(crypto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}

func (c *CryptoHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := chi.URLParam(r, "id")

		id, err := uuid.Parse(idString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := c.Store.DeleteCrypto(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
