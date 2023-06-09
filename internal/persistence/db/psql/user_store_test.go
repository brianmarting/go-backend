package psql

import (
	"errors"
	"go-backend/internal/persistence/db/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

var user = model.User{
	Id:   1,
	Uuid: uuid.New(),
	Name: "John",
}

func TestUserStore_GetByUuid(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &userStore{
		DB: sqlxDB,
	}

	tests := []struct {
		name    string
		args    model.User
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    user,
			wantErr: false,
			mock: func() {
				rows := sqlmock.
					NewRows([]string{"id", "uuid", "name"}).
					AddRow(user.Id, user.Uuid, user.Name)
				mock.ExpectQuery("SELECT (.+) FROM exchange_user WHERE uuid = (.+)").
					WithArgs(user.Uuid.String()).
					WillReturnRows(rows)
			},
		},
		{
			name:    "Should execute query with error",
			args:    user,
			wantErr: true,
			mock: func() {
				mock.ExpectQuery("SELECT (.+) FROM exchange_user WHERE uuid = (.+)").WithArgs(user.Uuid, user.Name).
					WithArgs(user.Uuid.String()).
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

			if err == nil {
				assert.Equal(t, result, user, "should be equal")
			}
		})
	}
}

func TestUserStore_Create(t *testing.T) {
	db, mock, _ := sqlmock.New()
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	s := &userStore{
		DB: sqlxDB,
	}

	tests := []struct {
		name    string
		args    model.User
		mock    func()
		wantErr bool
	}{
		{
			name:    "Should execute query without error",
			args:    user,
			wantErr: false,
			mock: func() {
				mock.ExpectExec("INSERT INTO exchange_user").WithArgs(user.Uuid, user.Name).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:    "Should execute query with error",
			args:    user,
			wantErr: true,
			mock: func() {
				mock.ExpectExec("INSERT INTO exchange_user").WithArgs(user.Uuid, user.Name).WillReturnError(errors.New("failed to exec query"))
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
