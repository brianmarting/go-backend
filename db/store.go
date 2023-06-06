package db

import (
	"fmt"
	"go-backend/db/psql"
	"sync"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var once sync.Once

var (
	store    *Store
	storeErr error
)

func GetStore() *Store {
	if store == nil {
		once.Do(func() {
			store, storeErr = newStore()
			if storeErr != nil {
				log.Fatal().Err(storeErr)
			}
		})
	}

	return store
}

func newStore() (*Store, error) {
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
		UserStore:   &UserStore{DB: db},
	}, nil
}

type Store struct {
	*CryptoStore
	*WalletStore
	*UserStore
}
