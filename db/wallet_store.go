package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Wallet struct {
	Id      int       `db:"id"`
	Uuid    uuid.UUID `db:"uuid"`
	Address string    `db:"address"`
}

type WalletStore struct {
	*sqlx.DB
}

func (s *WalletStore) GetByUuid(id uuid.UUID) (Wallet, error) {
	var w Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE uuid = $1", id); err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) GetByAddress(address string) (Wallet, error) {
	var w Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE address = $1", address); err != nil {
		return Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) Create(w Wallet) error {
	if _, err := s.Exec("INSERT INTO wallet (uuid, address) VALUES ($1, $2) RETURNING *", w.Uuid, w.Address); err != nil {
		return err
	}

	return nil
}
