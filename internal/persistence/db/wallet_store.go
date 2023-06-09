package db

import (
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type WalletStore struct {
	*sqlx.DB
}

func (s *WalletStore) GetByUuid(id uuid.UUID) (model.Wallet, error) {
	var w model.Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE uuid = $1", id); err != nil {
		return model.Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) GetByAddress(address string) (model.Wallet, error) {
	var w model.Wallet

	if err := s.Get(&w, "SELECT * FROM wallet WHERE address = $1", address); err != nil {
		return model.Wallet{}, err
	}

	return w, nil
}

func (s *WalletStore) UpdateAmountById(id int, amount int) error {
	_, err := s.Exec("UPDATE wallet SET amount = $1 WHERE id = $2", amount, id)

	return err
}

func (s *WalletStore) Create(w model.Wallet) error {
	if _, err := s.Exec("INSERT INTO wallet (uuid, address) VALUES ($1, $2) RETURNING *", w.Uuid, w.Address); err != nil {
		return err
	}

	return nil
}
