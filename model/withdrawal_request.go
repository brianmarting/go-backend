package model

import "github.com/google/uuid"

type WithdrawRequest struct {
	CryptoId    uuid.UUID
	FromAddress string
	ToAddress   string
	Amount      int
}
