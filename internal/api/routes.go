package api

import (
	handlers2 "go-backend/internal/api/handlers"
	"go-backend/internal/facade/queue"
	"go-backend/internal/persistence/db"
	service2 "go-backend/internal/service"
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

func NewHandler(store *db.Store) Handler {
	return handler{
		Mux: chi.NewMux(),
		cryptoHandler: handlers2.NewCryptoHandler(
			service2.NewCryptoService(store.CryptoStore),
		),
		walletHandler: handlers2.NewWalletHandler(
			service2.NewWalletService(store.WalletStore),
		),
		withdrawalHandler: handlers2.NewWithdrawalHandler(
			queue.NewPublisher(),
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
