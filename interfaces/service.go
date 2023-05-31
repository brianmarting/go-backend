package interfaces

import "go-backend/model"

type WithdrawalService interface {
	Withdraw(wr model.WithdrawalRequest) error
}
