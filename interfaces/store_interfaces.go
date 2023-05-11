package interfaces

import (
	"github.com/google/uuid"
	"go-backend/db"
)

type CryptoStore interface {
	Crypto(id uuid.UUID) (db.Crypto, error)
	CreateCrypto(c *db.Crypto) error
	DeleteCrypto(id uuid.UUID) error
}

type Store interface {
	CryptoStore
}
