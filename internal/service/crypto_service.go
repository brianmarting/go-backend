package service

import (
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/model"

	"github.com/google/uuid"
)

type CryptoService interface {
	GetByUuid(id uuid.UUID) (model.Crypto, error)
	Create(c model.Crypto) error
	Delete(id uuid.UUID) error
}

type cryptoService struct {
	cryptoStore db.CryptoStore
}

func NewCryptoService(store db.CryptoStore) CryptoService {
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
