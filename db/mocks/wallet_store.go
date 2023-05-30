package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"go-backend/db"
)

type WalletStoreMock struct {
	mock.Mock
}

func (m *WalletStoreMock) GetByUuid(id uuid.UUID) (db.Wallet, error) {
	args := m.Called(id)
	return args.Get(0).(db.Wallet), args.Error(1)
}

func (m *WalletStoreMock) GetByAddress(address string) (db.Wallet, error) {
	args := m.Called(address)
	return args.Get(0).(db.Wallet), args.Error(1)
}

func (m *WalletStoreMock) Create(w db.Wallet) error {
	args := m.Called(w)
	return args.Error(0)
}
