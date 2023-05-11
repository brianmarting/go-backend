package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"go-backend/db"
	"go-backend/handler"
	"go-backend/interfaces"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*chi.Mux

	store interfaces.Store
}

func main() {
	store, err := db.NewStore("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := newHandler(store)
	_ = http.ListenAndServe(":8888", h)
}

func newHandler(store interfaces.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	cryptoHandler := handler.CryptoHandler{Store: store}

	h.Route("/crypto", func(router chi.Router) {
		router.Get("/{id}", cryptoHandler.Get())
		router.Post("/", cryptoHandler.Create())
		router.Post("/{id}/delete", cryptoHandler.Delete())
	})

	return h
}
