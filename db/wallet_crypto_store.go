package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WalletCrypto struct {
	Id       uuid.UUID `db:"id"`
	WalletId uuid.UUID `db:"wallet_id"`
	CryptoId uuid.UUID `db:"crypto_id"`
	Amount   int       `db:"amount"`
}

type WalletCryptoStore struct {
	*sqlx.DB
}

func (s *WalletCryptoStore) FindByWalletIdAndCryptoId(walletId uuid.UUID, cryptoId uuid.UUID) (WalletCrypto, error) {
	var wc WalletCrypto

	if err := s.Get(&wc, "SELECT * FROM wallet_crypto WHERE wallet_id = $1 AND crypto_id = $2", walletId, cryptoId); err != nil {
		return WalletCrypto{}, err
	}

	return wc, nil
}

func (s *WalletCryptoStore) SetAmountByWalletId(walletId uuid.UUID, amount int) error {
	_, err := s.Exec("UPDATE wallet_crypto SET amount = $1 WHERE wallet_id = $2", amount, walletId)

	return err
}
