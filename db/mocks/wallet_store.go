package mocks

import (
	"go-backend/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type WalletStoreMock struct {
	mock.Mock
}

func (m *WalletStoreMock) GetByUuid(id uuid.UUID) (model.Wallet, error) {
	args := m.Called(id)
	return args.Get(0).(model.Wallet), args.Error(1)
}

func (m *WalletStoreMock) GetByAddress(address string) (model.Wallet, error) {
	args := m.Called(address)
	return args.Get(0).(model.Wallet), args.Error(1)
}

func (m *WalletStoreMock) UpdateAmountById(id int, amount int) error {
	args := m.Called(id, amount)
	return args.Error(0)
}

func (m *WalletStoreMock) Create(w model.Wallet) error {
	args := m.Called(w)
	return args.Error(0)
}
