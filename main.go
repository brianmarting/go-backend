package main

import (
	"go-backend/api"
	"go-backend/app/queue"
	queue2 "go-backend/facade/queue"
	"go-backend/persistence/db"
	"go-backend/service"
	"net/http"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store := db.GetStore()

	consumer := queue.NewWithdrawalConsumer(
		queue2.NewConsumer(),
		service.NewWithdrawalService(service.NewWalletService(store.WalletStore)),
	)
	consumer.StartConsuming()

	h := api.NewHandler(store)
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}
