package service

import (
	"go-backend/internal/api/model"
)

type WithdrawalService interface {
	Withdraw(wr *model.WithdrawalRequest) error
}

type withdrawalService struct {
	walletService WalletService
}

func NewWithdrawalService(walletService WalletService) WithdrawalService {
	return withdrawalService{
		walletService: walletService,
	}
}

func (s withdrawalService) Withdraw(wr *model.WithdrawalRequest) error {
	walletFrom, err := s.walletService.GetByAddress(wr.FromAddress)
	if err != nil {
		return err
	}

	walletTo, err := s.walletService.GetByAddress(wr.ToAddress)
	if err != nil {
		return err
	}

	if err = s.walletService.UpdateAmountById(walletFrom.Id, walletFrom.Amount-wr.Amount); err != nil {
		return err
	}

	if err = s.walletService.UpdateAmountById(walletTo.Id, walletTo.Amount+wr.Amount); err != nil {
		return err
	}

	return nil
}
