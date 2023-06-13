package main

import (
	"context"
	"fmt"
	"go-backend/internal/api"
	"go-backend/internal/app/queue"
	"go-backend/internal/app/socket"
	facadeDB "go-backend/internal/facade/db"
	facadeQueue "go-backend/internal/facade/queue"
	facadeSocket "go-backend/internal/facade/socket"
	grpc "go-backend/internal/grpc"
	pb "go-backend/internal/grpc/generated"
	"go-backend/internal/observability/tracing"
	"go-backend/internal/service"
	googleGrpc "google.golang.org/grpc"
	"net"
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

	go startGrpcServer(withdrawalService)

	h := api.NewHandler()
	_ = h.CreateAllRoutes()
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Info().Err(err).Msg("failed to listen and serve http handler")
	}
}

func startGrpcServer(service service.WithdrawalService) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start grpc server listener")
	}

	grpcServer := googleGrpc.NewServer()
	pb.RegisterWithdrawalServiceServer(grpcServer, grpc.NewGrpcWithdrawalServer(service))
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start grpc server")
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
