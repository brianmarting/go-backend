package db

import (
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
)

type WalletStore interface {
	GetByUuid(id uuid.UUID) (model.Wallet, error)
	GetByAddress(address string) (model.Wallet, error)
	UpdateAmountById(id int, amount int) error
	Create(w model.Wallet) error
}
