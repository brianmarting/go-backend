package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	Id      uuid.UUID `db:"id"`
	Address string    `db:"address"`
}

type WalletStore struct {
	*sqlx.DB
}

func (s *WalletStore) Wallet(id uuid.UUID) (Wallet, error) {
	var w Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE id = $1", id); err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) ByAddress(address string) (Wallet, error) {
	var w Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE address = $1", address); err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) CreateWallet(w *Wallet) error {
	if err := s.Get(&w, "INSERT INTO wallet VALUES ($1, $2) RETURNING *"); err != nil {
		return err
	}

	return nil
}
