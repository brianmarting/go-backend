package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewStore(driverName string) (*Store, error) {
	db, err := sqlx.Open("postgres", driverName)

	if err != nil {
		return nil, fmt.Errorf("Error opening db", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging db", err)
	}

	return &Store{
		CryptoStore:       &CryptoStore{DB: db},
		WalletStore:       &WalletStore{DB: db},
		WalletCryptoStore: &WalletCryptoStore{DB: db},
	}, nil
}

type Store struct {
	*CryptoStore
	*WalletStore
	*WalletCryptoStore
}
