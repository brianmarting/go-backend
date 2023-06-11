package model

import "github.com/google/uuid"

type WithdrawalRequest struct {
	CryptoId    uuid.UUID `json:"cryptoId"`
	FromAddress string    `json:"fromAddress"`
	ToAddress   string    `json:"toAddress"`
	Amount      int       `json:"amount"`
}
