package service

import (
	"fmt"
	"go-backend/interfaces"
	"go-backend/model"
)

type WithdrawalService struct {
	WalletStore       interfaces.WalletStore
	WalletCryptoStore interfaces.WalletCryptoStore
}

func (s *WithdrawalService) Withdraw(wr model.WithdrawRequest) error {
	walletFrom, err := s.WalletStore.ByAddress(wr.FromAddress)
	if err != nil {
		return err
	}

	walletCryptoFrom, err := s.WalletCryptoStore.FindByWalletIdAndCryptoId(walletFrom.Id, wr.CryptoId)
	if err != nil {
		return err
	}

	if walletCryptoFrom.Amount < wr.Amount {
		fmt.Println("You do not have a sufficient amount of tokens")
		return nil
	}

	s.WalletCryptoStore.SetAmountByWalletId(walletFrom.Id, walletCryptoFrom.Amount-wr.Amount)

	walletTo, err := s.WalletStore.ByAddress(wr.ToAddress)
	if err != nil {
		return err
	}

	walletCryptoTo, err := s.WalletCryptoStore.FindByWalletIdAndCryptoId(walletTo.Id, wr.CryptoId)
	if err != nil {
		return err
	}

	s.WalletCryptoStore.SetAmountByWalletId(walletTo.Id, walletCryptoTo.Amount+wr.Amount)

	return nil
}
