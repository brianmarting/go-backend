package api

import (
	handlers2 "go-backend/api/handlers"
	"go-backend/db"
	"go-backend/facade"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)

	CreateAllRoutes() Handler
}

type handler struct {
	*chi.Mux

	cryptoHandler     handlers2.CryptoHandler
	walletHandler     handlers2.WalletHandler
	withdrawalHandler handlers2.WithdrawalHandler
}

func NewHandler() Handler {
	store := db.GetStore()

	return handler{
		Mux:           chi.NewMux(),
		cryptoHandler: handlers2.NewCryptoHandler(store.CryptoStore),
		walletHandler: handlers2.NewWalletHandler(store.WalletStore),
		withdrawalHandler: handlers2.NewWithdrawalHandler(
			facade.GetPublisher(),
			store.CryptoStore,
			store.WalletStore,
		),
	}
}

func (h handler) CreateAllRoutes() Handler {
	h.Route("/crypto", func(router chi.Router) {
		router.Get("/{id}", h.cryptoHandler.Get())
		router.Post("/", h.cryptoHandler.Create())
		router.Post("/{id}/delete", h.cryptoHandler.Delete())

		router.Route("/wallet", func(router chi.Router) {
			router.Get("/{id}", h.walletHandler.Get())
			router.Post("/", h.walletHandler.Create())
		})

		router.Route("/withdraw", func(router chi.Router) {
			router.Post("/", h.withdrawalHandler.Withdraw())
		})
	})

	return h
}
