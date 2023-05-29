package db

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
)

var walletCrypto = WalletCrypto{
	Id:       1,
	Uuid:     uuid.New(),
	WalletId: 1,
	CryptoId: 2,
	Amount:   15,
}

func TestWalletCryptoStore_FindByWalletIdAndCryptoId(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &WalletCryptoStore{
		DB: sqlxDB,
	}
	type args struct {
		walletId int
		cryptoId int
	}
	tests := []struct {
		name    string
		args    args
		mock    func()
		wantErr bool
	}{
		{
			name: "Should execute query without error",
			args: args{
				walletId: 1,
				cryptoId: 2,
			},
			mock: func() {
				rows := sqlmock.
					NewRows([]string{"id", "uuid", "wallet_id", "crypto_id", "amount"}).
					AddRow(walletCrypto.Id, walletCrypto.Uuid, walletCrypto.WalletId, walletCrypto.CryptoId, walletCrypto.Amount)
				mock.ExpectQuery("SELECT (.+) FROM wallet_crypto WHERE wallet_id = (.+) AND crypto_id = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(walletCrypto.WalletId, walletCrypto.CryptoId).
					WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "Should execute query with error",
			args: args{
				walletId: 1,
				cryptoId: 2,
			},
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM wallet_crypto WHERE wallet_id = (.+) AND crypto_id = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(walletCrypto.WalletId, walletCrypto.CryptoId).
					WillReturnError(errors.New("failed to exec query"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				tt.mock()

				result, err := s.FindByWalletIdAndCryptoId(tt.args.walletId, tt.args.cryptoId)

				if (err != nil) != tt.wantErr {
					t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				if err == nil {
					assert.Equal(t, result, walletCrypto, "should be equal")
				}
			})
		})
	}
}
