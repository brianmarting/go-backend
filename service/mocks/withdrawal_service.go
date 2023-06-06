package mocks

import (
	"go-backend/model"

	"github.com/stretchr/testify/mock"
)

type WithdrawalServiceMock struct {
	mock.Mock
}

func (m *WithdrawalServiceMock) Withdraw(wr model.WithdrawalRequest) error {
	args := m.Called(wr)
	return args.Error(0)
}
