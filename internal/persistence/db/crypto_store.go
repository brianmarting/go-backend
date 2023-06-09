package db

import (
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
)

type CryptoStore interface {
	GetByUuid(id uuid.UUID) (model.Crypto, error)
	Create(c model.Crypto) error
	Delete(id uuid.UUID) error
}
