package service

import (
	"go-backend/persistence/db/model"

	"github.com/google/uuid"
)

type WalletStore interface {
	GetByUuid(id uuid.UUID) (model.Wallet, error)
	GetByAddress(address string) (model.Wallet, error)
	UpdateAmountById(id int, amount int) error
	Create(w model.Wallet) error
}

type WalletService interface {
	GetByUuid(id uuid.UUID) (model.Wallet, error)
	GetByAddress(address string) (model.Wallet, error)
	UpdateAmountById(id int, amount int) error
	Create(w model.Wallet) error
}

type walletService struct {
	walletStore WalletStore
}

func NewWalletService(store WalletStore) WalletService {
	return walletService{
		walletStore: store,
	}
}

func (s walletService) GetByUuid(id uuid.UUID) (model.Wallet, error) {
	return s.walletStore.GetByUuid(id)
}

func (s walletService) GetByAddress(address string) (model.Wallet, error) {
	return s.walletStore.GetByAddress(address)
}

func (s walletService) UpdateAmountById(id int, amount int) error {
	return s.walletStore.UpdateAmountById(id, amount)
}

func (s walletService) Create(w model.Wallet) error {
	return s.walletStore.Create(w)
}
