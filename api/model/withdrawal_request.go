package model

import "github.com/google/uuid"

type WithdrawalRequest struct {
	CryptoId    uuid.UUID
	FromAddress string
	ToAddress   string
	Amount      int
}
