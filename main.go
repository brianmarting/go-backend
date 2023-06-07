package main

import (
	"go-backend/api"
	"go-backend/app/queue"
	facadeQueue "go-backend/facade/queue"
	"go-backend/persistence/db"
	"go-backend/service"
	"net/http"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store := db.GetStore()

	consumer := createWithdrawalConsumer(store)
	consumer.StartConsuming()

	h := api.NewHandler(store)
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}

func createWithdrawalConsumer(store *db.Store) queue.WithdrawalConsumer {
	return queue.NewWithdrawalConsumer(
		facadeQueue.NewConsumer(),
		service.NewWithdrawalService(service.NewWalletService(store.WalletStore)),
	)
}
