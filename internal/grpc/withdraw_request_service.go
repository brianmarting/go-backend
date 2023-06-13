package grpc

import (
	"context"
	"go-backend/internal/api/model"
	pb "go-backend/internal/grpc/generated"
	"go-backend/internal/observability/tracing"
	"go-backend/internal/service"

	"github.com/google/uuid"
)

type WithdrawalServer struct {
	pb.UnimplementedWithdrawalServiceServer

	withdrawalService service.WithdrawalService
}

func NewGrpcWithdrawalServer(service service.WithdrawalService) WithdrawalServer {
	return WithdrawalServer{
		withdrawalService: service,
	}
}

func (w WithdrawalServer) SendStreaming(stream pb.WithdrawalService_SendStreamingServer) error {
	for {
		recv, err := stream.Recv()

		tracer := tracing.GetTracer()
		_, span := tracer.Start(context.Background(), "receive-withdrawal-msg-grpc-stream")
		defer span.End()

		if err != nil {
			return stream.SendAndClose(&pb.WithdrawRequestResult{
				Result: "NOK",
			})
		}

		wr, err := convertToWithdrawalRequest(recv)
		if err != nil {
			return stream.SendAndClose(&pb.WithdrawRequestResult{
				Result: "NOK",
			})
		}

		err = w.withdrawalService.Withdraw(&wr)
		if err != nil {
			return stream.SendAndClose(&pb.WithdrawRequestResult{
				Result: "NOK",
			})
		}
	}
}

func convertToWithdrawalRequest(wr *pb.WithdrawRequest) (model.WithdrawalRequest, error) {
	parsed, err := uuid.Parse(wr.GetCryptoId().Value)
	if err != nil {
		return model.WithdrawalRequest{}, err
	}

	return model.WithdrawalRequest{
		CryptoId:    parsed,
		FromAddress: wr.FromAddress,
		ToAddress:   wr.ToAddress,
		Amount:      int(wr.Amount),
	}, nil
}
