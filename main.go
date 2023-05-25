package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"go-backend/db"
	"go-backend/goroutines"
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

	// Create handler and listen+serve for requests in a blocking manner
	h := newHandler(store)
	_ = http.ListenAndServe(":8888", h)
}

func newHandler(store interfaces.Store) *Handler {
	h := &Handler{
		Mux:   chi.NewMux(),
		store: store,
	}

	h.Use(middleware.Logger)

	withdrawalHandler := handler.WithdrawalHandler{
		WalletStore:       store,
		WalletCryptoStore: store,
	}
	goroutines.StartDispatcher(1)

	cryptoHandler := handler.CryptoHandler{Store: store}
	walletHandler := handler.WalletHandler{Store: store}

	h.Route("/crypto", func(router chi.Router) {
		router.Get("/{id}", cryptoHandler.Get())
		router.Post("/", cryptoHandler.Create())
		router.Post("/{id}/delete", cryptoHandler.Delete())
	})

	h.Route("/wallet", func(router chi.Router) {
		router.Get("/{id}", walletHandler.Get())
		router.Post("/", walletHandler.Create())
	})

	h.Route("/withdraw", func(router chi.Router) {
		router.Post("/", withdrawalHandler.Withdraw())
	})

	return h
}
