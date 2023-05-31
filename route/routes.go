package route

import (
	"github.com/go-chi/chi/v5"
	"go-backend/handler"
)

type Handler struct {
	*chi.Mux

	cryptoHandler     *handler.CryptoHandler     `di.inject:"cryptoHandler"`
	walletHandler     *handler.WalletHandler     `di.inject:"walletHandler"`
	withdrawalHandler *handler.WithdrawalHandler `di.inject:"withdrawalHandler"`
}

func (h *Handler) CreateAllRoutes() *Handler {
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
