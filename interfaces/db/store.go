package db

import (
	"github.com/google/uuid"
	"go-backend/model"
)

type CryptoStore interface {
	GetByUuid(id uuid.UUID) (model.Crypto, error)
	Create(c model.Crypto) error
	Delete(id uuid.UUID) error
}

type WalletStore interface {
	GetByUuid(id uuid.UUID) (model.Wallet, error)
	GetByAddress(address string) (model.Wallet, error)
	UpdateAmountById(id int, amount int) error
	Create(w model.Wallet) error
}

type UserStore interface {
	GetByUuid(id uuid.UUID) (model.User, error)
	Create(u model.User) error
}
