package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewStore(driverName string) (*Store, error) {
	db, err := sqlx.Open("postgres", driverName)

	if err != nil {
		return nil, fmt.Errorf("Error opening db")
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Error pinging db")
	}

	return &Store{
		CryptoStore: &CryptoStore{DB: db},
	}, nil
}

type Store struct {
	*CryptoStore
}
