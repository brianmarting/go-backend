package route

import (
	"github.com/go-chi/chi/v5"
	"go-backend/db"
	"go-backend/facade"
	"go-backend/handler"
)

type Handler struct {
	*chi.Mux

	*handler.CryptoHandler
	*handler.WalletHandler
	*handler.WithdrawalHandler
}

func NewHandler() *Handler {
	store := db.GetStore()
	return &Handler{
		Mux:           chi.NewMux(),
		CryptoHandler: &handler.CryptoHandler{Store: store.CryptoStore},
		WalletHandler: &handler.WalletHandler{Store: store.WalletStore},
		WithdrawalHandler: &handler.WithdrawalHandler{
			Publisher:   facade.GetPublisher(),
			WalletStore: store.WalletStore,
			CryptoStore: store.CryptoStore,
		},
	}
}

func (h *Handler) CreateAllRoutes() *Handler {
	h.Route("/crypto", func(router chi.Router) {
		router.Get("/{id}", h.CryptoHandler.Get())
		router.Post("/", h.CryptoHandler.Create())
		router.Post("/{id}/delete", h.CryptoHandler.Delete())

		router.Route("/wallet", func(router chi.Router) {
			router.Get("/{id}", h.WalletHandler.Get())
			router.Post("/", h.WalletHandler.Create())
		})

		router.Route("/withdraw", func(router chi.Router) {
			router.Post("/", h.WithdrawalHandler.Withdraw())
		})
	})

	return h
}
