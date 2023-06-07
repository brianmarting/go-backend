package main

import (
	"go-backend/api"
	"go-backend/app/queue"
	"go-backend/app/socket"
	facadeQueue "go-backend/facade/queue"
	facadeSocket "go-backend/facade/socket"
	"go-backend/persistence/db"
	"go-backend/service"
	"net/http"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store := db.GetStore()

	withdrawalService := service.NewWithdrawalService(
		service.NewWalletService(store.WalletStore),
	)

	consumer := createWithdrawalConsumer(withdrawalService)
	consumer.Start()

	withdrawalSocketListener := createWithdrawalSocketListener(withdrawalService)
	withdrawalSocketListener.Start()

	h := api.NewHandler(store)
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}

func createWithdrawalConsumer(withdrawalService service.WithdrawalService) queue.WithdrawalConsumer {
	return queue.NewWithdrawalConsumer(
		facadeQueue.NewConsumer(),
		withdrawalService,
	)
}

func createWithdrawalSocketListener(withdrawalService service.WithdrawalService) socket.WithdrawalSocketListener {
	return socket.NewWithdrawalSocketListener(
		facadeSocket.NewTcpSocketListener(),
		withdrawalService,
	)
}
