package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Crypto struct {
	Id          int       `db:"id"`
	Uuid        uuid.UUID `db:"uuid"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type CryptoStore struct {
	*sqlx.DB
}

func (s *CryptoStore) GetByUuid(id uuid.UUID) (Crypto, error) {
	var c Crypto

	if err := s.Get(&c, "SELECT * FROM crypto WHERE uuid = $1", id.String()); err != nil {
		return Crypto{}, err
	}

	return c, nil
}

func (s *CryptoStore) Create(c Crypto) error {
	_, err := s.Exec("INSERT INTO crypto (uuid, name, description) VALUES ($1, $2, $3)", c.Uuid, c.Name, c.Description)

	return err
}

func (s *CryptoStore) Delete(id uuid.UUID) error {
	if _, err := s.Exec("DELETE FROM crypto WHERE uuid = $1", id); err != nil {
		return err
	}

	return nil
}
