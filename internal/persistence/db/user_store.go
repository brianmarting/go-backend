package db

import (
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
)

type UserStore interface {
	GetByUuid(id uuid.UUID) (model.User, error)
	Create(u model.User) error
}
