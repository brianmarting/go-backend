package worker

import "github.com/stretchr/testify/mock"

type WithdrawalWorkerJobMock struct {
	mock.Mock
}

func (w *WithdrawalWorkerJobMock) Work() error {
	args := w.Called()
	return args.Error(0)
}
