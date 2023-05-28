package interfaces

import (
	"github.com/google/uuid"
	"go-backend/db"
)

type CryptoStore interface {
	GetByUuid(id uuid.UUID) (db.Crypto, error)
	Create(c db.Crypto) error
	Delete(id uuid.UUID) error
}

type WalletStore interface {
	GetByUuid(id uuid.UUID) (db.Wallet, error)
	GetByAddress(address string) (db.Wallet, error)
	Create(w db.Wallet) error
}

type WalletCryptoStore interface {
	FindByWalletIdAndCryptoId(walletId int, cryptoId int) (db.WalletCrypto, error)
	SetAmountByWalletId(walletId int, amount int) error
}
