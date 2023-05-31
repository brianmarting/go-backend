package db

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id   int       `db:"id"`
	Uuid uuid.UUID `db:"uuid"`
	Name string    `db:"name"`
}

type UserStore struct {
	*sqlx.DB
}

func (s *UserStore) GetByUuid(id uuid.UUID) (User, error) {
	var u User

	if err := s.Get(&u, "SELECT * FROM exchange_user WHERE uuid = $1", id.String()); err != nil {
		return User{}, err
	}

	return u, nil
}

func (s *UserStore) Create(u User) error {
	_, err := s.Exec("INSERT INTO exchange_user (uuid, name) VALUES ($1, $2)", u.Uuid, u.Name)

	return err
}
