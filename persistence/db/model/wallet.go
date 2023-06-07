package model

import "github.com/google/uuid"

type Wallet struct {
	Id       int       `db:"id"`
	Uuid     uuid.UUID `db:"uuid"`
	CryptoId int       `db:"crypto_id"`
	Address  string    `db:"address"`
	Amount   int       `db:"amount"`
}
