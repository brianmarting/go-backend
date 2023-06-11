package service

import (
	"go-backend/internal/persistence/db"
	"go-backend/internal/persistence/db/mocks"
	"go-backend/internal/persistence/db/model"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

func Test_cryptoService_Create(t *testing.T) {
	crypto := model.Crypto{}
	cryptoStoreMock := new(mocks.CryptoStoreMock)
	service := &cryptoService{
		cryptoStore: cryptoStoreMock,
	}
	tests := []struct {
		name        string
		args        model.Crypto
		mockFn      func() *mock.Call
		expectError bool
	}{
		{
			name: "should create",
			args: crypto,
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("Create", crypto).Return(nil)
			},
		},
		//{
		//	name: "should return err",
		//	args: crypto,
		//	mockFn: func() *mock.Call {
		//		return cryptoStoreMock.On("Create", crypto).Return(errors.New("failed"))
		//	},
		//	expectError: true,
		//},
	}
	for _, tt := range tests {
		var mockCall *mock.Call
		if tt.mockFn != nil {
			mockCall = tt.mockFn()
		}
		defer func() {
			if mockCall != nil {
				mockCall.Unset()
			}
		}()

		t.Run(tt.name, func(t *testing.T) {
			if err := service.Create(tt.args); (err != nil) != tt.expectError {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.expectError)
			}
		})
	}
}

func Test_cryptoService_Delete(t *testing.T) {
	type fields struct {
		cryptoStore db.CryptoStore
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cryptoService{
				cryptoStore: tt.fields.cryptoStore,
			}
			if err := s.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_cryptoService_GetByUuid(t *testing.T) {
	type fields struct {
		cryptoStore db.CryptoStore
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Crypto
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := cryptoService{
				cryptoStore: tt.fields.cryptoStore,
			}
			got, err := s.GetByUuid(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByUuid() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUuid() got = %v, want %v", got, tt.want)
			}
		})
	}
}
