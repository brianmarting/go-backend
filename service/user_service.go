package service

import (
	"github.com/google/uuid"
	"go-backend/persistence/db/model"
)

type UserStore interface {
	GetByUuid(id uuid.UUID) (model.User, error)
	Create(u model.User) error
}

type UserService interface {
	GetByUuid(id uuid.UUID) (model.User, error)
	Create(u model.User) error
}

type userService struct {
	store UserStore
}

func NewUserService(store UserStore) UserService {
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
