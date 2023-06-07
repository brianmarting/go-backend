package service

import (
	"github.com/google/uuid"
	"go-backend/persistence/db/model"
)

type CryptoStore interface {
	GetByUuid(id uuid.UUID) (model.Crypto, error)
	Create(c model.Crypto) error
	Delete(id uuid.UUID) error
}

type CryptoService interface {
	GetByUuid(id uuid.UUID) (model.Crypto, error)
	Create(c model.Crypto) error
	Delete(id uuid.UUID) error
}

type cryptoService struct {
	cryptoStore CryptoStore
}

func NewCryptoService(store CryptoStore) CryptoService {
	return cryptoService{
		cryptoStore: store,
	}
}

func (s cryptoService) GetByUuid(id uuid.UUID) (model.Crypto, error) {
	return s.cryptoStore.GetByUuid(id)
}

func (s cryptoService) Create(c model.Crypto) error {
	return s.cryptoStore.Create(c)
}

func (s cryptoService) Delete(id uuid.UUID) error {
	return s.cryptoStore.Delete(id)
}
