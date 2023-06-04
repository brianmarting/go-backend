package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"go-backend/model"
)

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) GetByUuid(id uuid.UUID) (model.User, error) {
	var u model.User

	if err := s.Get(&u, "SELECT * FROM exchange_user WHERE uuid = $1", id.String()); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (s *UserStore) Create(u model.User) error {
	_, err := s.Exec("INSERT INTO exchange_user (uuid, name) VALUES ($1, $2)", u.Uuid, u.Name)

	return err
}
