package db

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testing"
)

var wallet = Wallet{
	Id:      1,
	Uuid:    uuid.New(),
	Address: "024Xoefeof",
}

func TestWalletStore_GetByUuid(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &WalletStore{
		DB: sqlxDB,
	}

	tests := []struct {
		name    string
		args    Wallet
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    wallet,
			wantErr: false,
			mock: func() {
				rows := sqlmock.
					NewRows([]string{"id", "uuid", "address"}).
					AddRow(wallet.Id, wallet.Uuid, wallet.Address)
				mock.ExpectQuery("SELECT (.+) FROM wallet WHERE uuid = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(wallet.Uuid.String()).
					WillReturnRows(rows)
			},
		},
		{
			name:    "Should execute query with error",
			args:    wallet,
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM wallet WHERE uuid = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(crypto.Uuid.String()).
					WillReturnError(errors.New("failed to exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := s.GetByUuid(tt.args.Uuid)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result != wallet {
				t.Errorf("Bodies dont match")
			}
		})
	}
}

func TestWalletStore_GetByAddress(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &WalletStore{
		DB: sqlxDB,
	}

	tests := []struct {
		name    string
		args    Wallet
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    wallet,
			wantErr: false,
			mock: func() {
				rows := sqlmock.
					NewRows([]string{"id", "uuid", "address"}).
					AddRow(wallet.Id, wallet.Uuid, wallet.Address)
				mock.ExpectQuery("SELECT (.+) FROM wallet WHERE address = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(wallet.Address).
					WillReturnRows(rows)
			},
		},
		{
			name:    "Should execute query with error",
			args:    wallet,
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM wallet WHERE address = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(wallet.Address).
					WillReturnError(errors.New("failed to exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			result, err := s.GetByAddress(tt.args.Address)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetByAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result != wallet {
				t.Errorf("Bodies dont match")
			}
		})
	}
}

func TestWalletStore_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &WalletStore{
		DB: sqlxDB,
	}

	tests := []struct {
		name    string
		args    Wallet
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    wallet,
			wantErr: false,
			mock: func() {
				mock.ExpectExec("INSERT INTO wallet").WithArgs(wallet.Uuid, wallet.Address).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "Should execute query with error",
			args:    wallet,
			wantErr: true,
			mock: func() {
				mock.ExpectExec("INSERT INTO crypto").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).WillReturnError(errors.New("failed to exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			if err := s.Create(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
