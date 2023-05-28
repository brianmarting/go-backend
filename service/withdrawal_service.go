package service

import (
	"github.com/rs/zerolog/log"
	"go-backend/interfaces"
	"go-backend/model"
)

type WithdrawalService struct {
	CryptoStore       interfaces.CryptoStore
	WalletStore       interfaces.WalletStore
	WalletCryptoStore interfaces.WalletCryptoStore
}

func (s *WithdrawalService) Withdraw(wr model.WithdrawRequest) error {
	crypto, err := s.CryptoStore.Crypto(wr.CryptoId)
	if err != nil {
		return err
	}

	walletFrom, err := s.WalletStore.ByAddress(wr.FromAddress)
	if err != nil {
		return err
	}

	walletCryptoFrom, err := s.WalletCryptoStore.FindByWalletIdAndCryptoId(walletFrom.Id, crypto.Id)
	if err != nil {
		return err
	}

	if walletCryptoFrom.Amount < wr.Amount {
		log.Info().Msg("You do not have a sufficient amount of tokens")
		return nil
	}

	s.WalletCryptoStore.SetAmountByWalletId(walletFrom.Id, walletCryptoFrom.Amount-wr.Amount)

	walletTo, err := s.WalletStore.ByAddress(wr.ToAddress)
	if err != nil {
		return err
	}

	walletCryptoTo, err := s.WalletCryptoStore.FindByWalletIdAndCryptoId(walletTo.Id, crypto.Id)
	if err != nil {
		return err
	}

	s.WalletCryptoStore.SetAmountByWalletId(walletTo.Id, walletCryptoTo.Amount+wr.Amount)

	return nil
}
