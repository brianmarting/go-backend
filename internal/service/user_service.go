package service

import (
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
)

type UserService interface {
	GetByUuid(id uuid.UUID) (model.User, error)
	Create(u model.User) error
}

type userService struct {
	store db.UserStore
}

func NewUserService(store db.UserStore) UserService {
	return userService{
		store: store,
	}
}

func (s userService) GetByUuid(id uuid.UUID) (model.User, error) {
	return s.store.GetByUuid(id)
}

func (s userService) Create(u model.User) error {
	return s.store.Create(u)
}
