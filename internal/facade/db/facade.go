package db

import (
	"fmt"
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/psql"
	"os"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

var (
	once       = sync.Once{}
	dbInstance *sqlx.DB
	err        error
)

func NewCryptoStore() db.CryptoStore {
	return psql.NewCryptoStore(getDB())
}

func NewWalletStore() db.WalletStore {
	return psql.NewWalletStore(getDB())
}

func NewUserStore() db.UserStore {
	return psql.NewUserStore(getDB())
}

func getDB() *sqlx.DB {
	once.Do(func() {
		var (
			username = os.Getenv("DB_USERNAME")
			password = os.Getenv("DB_PASSWORD")
			host     = os.Getenv("DB_HOST")
		)

		dbInstance, err = sqlx.Open(
			"postgres",
			fmt.Sprintf("postgres://%s:%s@%s:5432/postgres?sslmode=disable", username, password, host),
		)
		if err != nil {
			log.Fatal().Err(err).Msg("failed to open connection to db")
		}
	})
	return dbInstance
}
