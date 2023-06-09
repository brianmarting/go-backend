package psql

import (
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type cryptoStore struct {
	*sqlx.DB
}

func NewCryptoStore(db *sqlx.DB) db.CryptoStore {
	return &cryptoStore{
		DB: db,
	}
}

func (s *cryptoStore) GetByUuid(id uuid.UUID) (model.Crypto, error) {
	var c model.Crypto

	if err := s.Get(&c, "SELECT * FROM crypto WHERE uuid = $1", id.String()); err != nil {
		return model.Crypto{}, err
	}

	return c, nil
}

func (s *cryptoStore) Create(c model.Crypto) error {
	_, err := s.Exec("INSERT INTO crypto (uuid, name, description) VALUES ($1, $2, $3)", c.Uuid, c.Name, c.Description)

	return err
}

func (s *cryptoStore) Delete(id uuid.UUID) error {
	if _, err := s.Exec("DELETE FROM crypto WHERE uuid = $1", id); err != nil {
		return err
	}

	return nil
}
