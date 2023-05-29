package mocks

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"go-backend/db"
)

type CryptoStoreMock struct {
	mock.Mock
}

func (m *CryptoStoreMock) GetByUuid(id uuid.UUID) (db.Crypto, error) {
	args := m.Called(id)
	return args.Get(0).(db.Crypto), args.Error(1)
}

func (m *CryptoStoreMock) Create(c db.Crypto) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *CryptoStoreMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
