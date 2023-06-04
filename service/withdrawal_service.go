package service

import (
	dbImpl "go-backend/db"
	"go-backend/interfaces/db"
	"go-backend/interfaces/service"
	"go-backend/model"
	"sync"
)

var once sync.Once

var withdrawalService *WithdrawalService

type WithdrawalService struct {
	db.WalletStore
}

func GetWithdrawalService() service.WithdrawalService {
	if withdrawalService == nil {
		once.Do(func() {
			store := dbImpl.GetStore()
			withdrawalService = &WithdrawalService{
				WalletStore: store.WalletStore,
			}
		})
	}

	return withdrawalService
}

func (s *WithdrawalService) Withdraw(wr *model.WithdrawalRequest) error {
	walletFrom, err := s.WalletStore.GetByAddress(wr.FromAddress)
	if err != nil {
		return err
	}

	walletTo, err := s.WalletStore.GetByAddress(wr.ToAddress)
	if err != nil {
		return err
	}

	if err = s.WalletStore.UpdateAmountById(walletFrom.Id, walletFrom.Amount-wr.Amount); err != nil {
		return err
	}

	if err = s.WalletStore.UpdateAmountById(walletTo.Id, walletTo.Amount+wr.Amount); err != nil {
		return err
	}

	return nil
}
