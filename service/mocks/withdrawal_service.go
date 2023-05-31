package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-backend/model"
)

type WithdrawalServiceMock struct {
	mock.Mock
}

func (m *WithdrawalServiceMock) Withdraw(wr model.WithdrawalRequest) error {
	args := m.Called(wr)
	return args.Error(0)
}
