package db

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestCryptoStore_CreateCrypto(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &CryptoStore{
		DB: sqlxDB,
	}

	crypto := Crypto{
		Id:          1,
		Uuid:        uuid.New(),
		Name:        "btc",
		Description: "Bitcoin",
	}

	tests := []struct {
		name    string
		args    Crypto
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    crypto,
			wantErr: false,
			mock: func() {
				mock.ExpectExec("INSERT INTO crypto").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "Should execute query with error",
			args:    crypto,
			wantErr: true,
			mock: func() {
				mock.ExpectExec("INSERT INTO crypto").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).WillReturnError(errors.New("failed to exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			if err := s.CreateCrypto(&tt.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateCrypto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
