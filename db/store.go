package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"go-backend/db/psql"
)

func NewStore() (*Store, error) {
	db, err := psql.NewDB()

	if err != nil {
		return nil, fmt.Errorf("error opening db %g", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging db %g", err)
	}

	return &Store{
		CryptoStore: &CryptoStore{DB: db},
		WalletStore: &WalletStore{DB: db},
	}, nil
}

type Store struct {
	*CryptoStore
	*WalletStore
}
