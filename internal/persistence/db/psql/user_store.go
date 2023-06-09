package psql

import (
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type userStore struct {
	*sqlx.DB
}

func NewUserStore(db *sqlx.DB) db.UserStore {
	return &userStore{
		DB: db,
	}
}

func (s *userStore) GetByUuid(id uuid.UUID) (model.User, error) {
	var u model.User

	if err := s.Get(&u, "SELECT * FROM exchange_user WHERE uuid = $1", id.String()); err != nil {
		return model.User{}, err
	}

	return u, nil
}

func (s *userStore) Create(u model.User) error {
	_, err := s.Exec("INSERT INTO exchange_user (uuid, name) VALUES ($1, $2)", u.Uuid, u.Name)

	return err
}
