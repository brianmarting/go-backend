package handlers

import (
	"encoding/json"
	"go-backend/internal/persistence/db/model"
	"go-backend/internal/service"
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
	service service.CryptoService
}

func NewCryptoHandler(service service.CryptoService) CryptoHandler {
	return cryptoHandler{
		service: service,
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

		crypto, err := c.service.GetByUuid(id)
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

		if err := c.service.Create(crypto); err != nil {
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

		if err := c.service.Delete(id); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(200)
	}
}
