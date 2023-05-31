package worker

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go-backend/model"
	"go-backend/service/mocks"
	"testing"
	"time"
)

func TestStartDispatcher(t *testing.T) {
	withdrawalServiceMock := new(mocks.WithdrawalServiceMock)

	t.Run("Should run withdraw request and withdraw", func(t *testing.T) {
		withdrawRequest := model.WithdrawalRequest{
			CryptoId:    uuid.UUID{},
			FromAddress: "X-FROM",
			ToAddress:   "X-TO",
			Amount:      10,
		}
		withdrawalServiceMock.On("Withdraw", withdrawRequest).Return(nil)

		StartDispatcher(1)
		WorkRequestChannel <- withdrawRequest

		assert.EventuallyWithT(t, func(c *assert.CollectT) {
			withdrawalServiceMock.AssertExpectations(t)
		}, 5*time.Second, time.Second)
	})
}
