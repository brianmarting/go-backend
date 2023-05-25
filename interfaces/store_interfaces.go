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

type WalletStore interface {
	Wallet(id uuid.UUID) (db.Wallet, error)
	ByAddress(address string) (db.Wallet, error)
	CreateWallet(w *db.Wallet) error
}

type WalletCryptoStore interface {
	FindByWalletIdAndCryptoId(walletId uuid.UUID, cryptoId uuid.UUID) (db.WalletCrypto, error)
	SetAmountByWalletId(walletId uuid.UUID, amount int) error
}

type Store interface {
	CryptoStore
	WalletStore
	WalletCryptoStore
}
