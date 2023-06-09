package main

import (
	"go-backend/internal/api"
	"go-backend/internal/app/queue"
	"go-backend/internal/app/socket"
	facadeQueue "go-backend/internal/facade/queue"
	facadeSocket "go-backend/internal/facade/socket"
	"go-backend/internal/persistence/db"
	service2 "go-backend/internal/service"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store := db.GetStore()

	withdrawalService := service2.NewWithdrawalService(
		service2.NewWalletService(store.WalletStore),
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

func createWithdrawalConsumer(withdrawalService service2.WithdrawalService) queue.WithdrawalConsumer {
	return queue.NewWithdrawalConsumer(
		facadeQueue.NewConsumer(),
		withdrawalService,
	)
}

func createWithdrawalSocketListener(withdrawalService service2.WithdrawalService) socket.WithdrawalSocketListener {
	port := os.Getenv("TCP_WITHDRAWAL_LISTENER_PORT")

	return socket.NewWithdrawalSocketListener(
		facadeSocket.NewTcpSocketListener(port),
		withdrawalService,
	)
}
