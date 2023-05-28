package main

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-backend/db"
	"go-backend/goroutines"
	"go-backend/handler"
	"go-backend/interfaces"
	"go-backend/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	*chi.Mux

	store interfaces.Store
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store, err := db.NewStore("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err)
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
	withdrawalService := service.WithdrawalService{
		CryptoStore:       store,
		WalletStore:       store,
		WalletCryptoStore: store,
	}
	goroutines.StartDispatcher(1, withdrawalService)

	cryptoHandler := handler.CryptoHandler{Store: store}
	walletHandler := handler.WalletHandler{Store: store}

	h.Route("/crypto", func(router chi.Router) {
		router.Get("/{id}", cryptoHandler.Get())
		router.Post("/", cryptoHandler.Create())
		router.Post("/{id}/delete", cryptoHandler.Delete())

		router.Route("/wallet", func(router chi.Router) {
			router.Get("/{id}", walletHandler.Get())
			router.Post("/", walletHandler.Create())
		})

		router.Route("/withdraw", func(router chi.Router) {
			router.Post("/", withdrawalHandler.Withdraw())
		})
	})

	return h
}
