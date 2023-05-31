package service

import "go-backend/db"

func NewWithdrawalService(store *db.Store) *WithdrawalService {
	return &WithdrawalService{
		CryptoStore: store.CryptoStore,
		WalletStore: store.WalletStore,
	}
}
