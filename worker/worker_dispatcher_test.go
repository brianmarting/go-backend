package worker

import (
	"github.com/stretchr/testify/assert"
	"go-backend/service/mocks"
	worker "go-backend/worker/mocks"
	"testing"
	"time"
)

func TestStartDispatcher(t *testing.T) {
	withdrawalServiceMock := new(mocks.WithdrawalServiceMock)

	t.Run("Should run work", func(t *testing.T) {
		withdrawalWorkerJobMock := new(worker.WithdrawalWorkerJobMock)

		withdrawalWorkerJobMock.On("Work").Return(nil)
		StartDispatcher(1)
		WorkRequestChannel <- withdrawalWorkerJobMock

		assert.EventuallyWithT(t, func(c *assert.CollectT) {
			withdrawalServiceMock.AssertExpectations(t)
		}, 5*time.Second, time.Second)
	})
}
