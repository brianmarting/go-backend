package db

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"testing"
)

var crypto = Crypto{
	Id:          1,
	Uuid:        uuid.New(),
	Name:        "btc",
	Description: "Bitcoin",
}

func TestCryptoStore_GetByUuid(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &CryptoStore{
		DB: sqlxDB,
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
				rows := sqlmock.
					NewRows([]string{"id", "uuid", "name", "description"}).
					AddRow(crypto.Id, crypto.Uuid, crypto.Name, crypto.Description)
				mock.ExpectQuery("SELECT (.+) FROM crypto WHERE uuid = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
					WithArgs(crypto.Uuid.String()).
					WillReturnRows(rows)
			},
		},
		{
			name:    "Should execute query with error",
			args:    crypto,
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM crypto WHERE uuid = (.+)").WithArgs(crypto.Uuid, crypto.Name, crypto.Description).
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
				t.Errorf("CreateCrypto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err == nil && result != crypto {
				t.Errorf("Bodies dont match")
			}
		})
	}
}

func TestCryptoStore_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &CryptoStore{
		DB: sqlxDB,
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

			if err := s.Create(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCryptoStore_DeleteCrypto(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &CryptoStore{
		DB: sqlxDB,
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
				mock.ExpectExec("DELETE FROM crypto WHERE uuid = (.+)").
					WithArgs(crypto.Uuid).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "Should execute query with error",
			args:    crypto,
			wantErr: true,
			mock: func() {
				mock.ExpectExec("DELETE FROM crypto WHERE uuid = (.+)").
					WithArgs(crypto.Uuid).
					WillReturnError(errors.New("failed to exec query"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := s.Delete(tt.args.Uuid)

			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
