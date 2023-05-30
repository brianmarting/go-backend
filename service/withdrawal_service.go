package service

import (
	"github.com/rs/zerolog/log"
	"go-backend/db"
	"go-backend/model"
)

type WithdrawalService struct {
	cryptoStore       *db.CryptoStore       `di.inject:"cryptoStore"`
	walletStore       *db.WalletStore       `di.inject:"walletStore"`
	walletCryptoStore *db.WalletCryptoStore `di.inject:"walletCryptoStore"`
}

func (s *WithdrawalService) Withdraw(wr model.WithdrawalRequest) error {
	crypto, err := s.cryptoStore.GetByUuid(wr.CryptoId)
	if err != nil {
		return err
	}

	walletFrom, err := s.walletStore.GetByAddress(wr.FromAddress)
	if err != nil {
		return err
	}

	walletCryptoFrom, err := s.walletCryptoStore.FindByWalletIdAndCryptoId(walletFrom.Id, crypto.Id)
	if err != nil {
		return err
	}

	if walletCryptoFrom.Amount < wr.Amount {
		log.Info().Msg("You do not have a sufficient amount of tokens")
		return nil
	}

	s.walletCryptoStore.SetAmountByWalletId(walletFrom.Id, walletCryptoFrom.Amount-wr.Amount)

	walletTo, err := s.walletStore.GetByAddress(wr.ToAddress)
	if err != nil {
		return err
	}

	walletCryptoTo, err := s.walletCryptoStore.FindByWalletIdAndCryptoId(walletTo.Id, crypto.Id)
	if err != nil {
		return err
	}

	s.walletCryptoStore.SetAmountByWalletId(walletTo.Id, walletCryptoTo.Amount+wr.Amount)

	return nil
}
