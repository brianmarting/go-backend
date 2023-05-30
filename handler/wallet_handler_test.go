package handler

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-backend/db"
	"go-backend/db/mocks"
	"net/http"
	"net/http/httptest"
	"testing"
)

var wallet = db.Wallet{}

func TestWalletHandler_Get(t *testing.T) {
	walletStoreMock := new(mocks.WalletStoreMock)
	w := &WalletHandler{
		Store: walletStoreMock,
	}
	uuid := uuid.New()

	tests := []struct {
		name         string
		rec          *httptest.ResponseRecorder
		reqFn        func() *http.Request
		mockFn       func() *mock.Call
		expectResult bool
		expectError  bool
	}{
		{
			name: "Should return err if url param is no uuid",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", "malformed")

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			expectError: true,
		},
		{
			name: "Should return err if db returns err",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", uuid.String())

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			mockFn: func() *mock.Call {
				return walletStoreMock.On("GetByUuid", uuid).Return(db.Wallet{}, errors.New("db err"))
			},
			expectError: true,
		},
		{
			name: "Should return wallet",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", uuid.String())

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			mockFn: func() *mock.Call {
				return walletStoreMock.On("GetByUuid", uuid).Return(wallet, nil)
			},
			expectResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockCall *mock.Call
			if tt.mockFn != nil {
				mockCall = tt.mockFn()
			}
			defer func() {
				if mockCall != nil {
					mockCall.Unset()
				}
			}()

			fn := w.Get()
			fn(tt.rec, tt.reqFn())

			if tt.expectError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				var result db.Wallet
				json.NewDecoder(tt.rec.Body).Decode(&result)

				assert.Equal(t, wallet, result)
				assert.Equal(t, "application/json", tt.rec.Header().Get("Content-Type"))
			}
			walletStoreMock.AssertExpectations(t)
		})
	}
}

func TestWalletHandler_Create(t *testing.T) {
	walletStoreMock := new(mocks.WalletStoreMock)
	w := &WalletHandler{
		Store: walletStoreMock,
	}

	tests := []struct {
		name         string
		rec          *httptest.ResponseRecorder
		reqFn        func() *http.Request
		mockFn       func() *mock.Call
		expectResult bool
		expectError  bool
	}{
		{
			name: "Should return err if db returns err",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				return httptest.NewRequest("GET", "/", nil)
			},
			mockFn: func() *mock.Call {
				return walletStoreMock.On("Create", mock.Anything).Return(errors.New("db err"))
			},
			expectError: true,
		},
		{
			name: "Should create wallet",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				return httptest.NewRequest("GET", "/", nil)
			},
			mockFn: func() *mock.Call {
				return walletStoreMock.On("Create", mock.Anything).Return(nil)
			},
			expectResult: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockCall *mock.Call
			if tt.mockFn != nil {
				mockCall = tt.mockFn()
			}
			defer func() {
				if mockCall != nil {
					mockCall.Unset()
				}
			}()

			fn := w.Create()
			fn(tt.rec, tt.reqFn())

			if tt.expectError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				assert.Equal(t, http.StatusOK, tt.rec.Code)
			}
			walletStoreMock.AssertExpectations(t)
		})
	}
}