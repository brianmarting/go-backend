package mocks

import (
	"go-backend/model"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type CryptoStoreMock struct {
	mock.Mock
}

func (m *CryptoStoreMock) GetByUuid(id uuid.UUID) (model.Crypto, error) {
	args := m.Called(id)
	return args.Get(0).(model.Crypto), args.Error(1)
}

func (m *CryptoStoreMock) Create(c model.Crypto) error {
	args := m.Called(c)
	return args.Error(0)
}

func (m *CryptoStoreMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}
