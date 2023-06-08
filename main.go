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
	"os"

	"github.com/rs/zerolog/log"

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
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Info().Err(err).Msg("failed to listen and serve http handler")
	}
}

func createWithdrawalConsumer(withdrawalService service.WithdrawalService) queue.WithdrawalConsumer {
	return queue.NewWithdrawalConsumer(
		facadeQueue.NewConsumer(),
		withdrawalService,
	)
}

func createWithdrawalSocketListener(withdrawalService service.WithdrawalService) socket.WithdrawalSocketListener {
	port := os.Getenv("TCP_WITHDRAWAL_LISTENER_PORT")

	return socket.NewWithdrawalSocketListener(
		facadeSocket.NewTcpSocketListener(port),
		withdrawalService,
	)
}
