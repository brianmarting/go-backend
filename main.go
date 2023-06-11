package main

import (
	"context"
	"go-backend/internal/api"
	"go-backend/internal/app/queue"
	"go-backend/internal/app/socket"
	facadeDB "go-backend/internal/facade/db"
	facadeQueue "go-backend/internal/facade/queue"
	facadeSocket "go-backend/internal/facade/socket"
	"go-backend/internal/observability/tracing"
	"go-backend/internal/service"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	tracer := tracing.InitTracerProvider()
	defer func() {
		if err := tracer.Shutdown(context.Background()); err != nil {
			log.Info().Err(err).Msg("failed to shut down tracer provider")
		}
	}()

	withdrawalService := service.NewWithdrawalService(
		service.NewWalletService(facadeDB.NewWalletStore()),
	)

	consumer := createWithdrawalConsumer(withdrawalService)
	consumer.Start()

	withdrawalSocketListener := createWithdrawalSocketListener(withdrawalService)
	withdrawalSocketListener.Start()

	h := api.NewHandler()
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
