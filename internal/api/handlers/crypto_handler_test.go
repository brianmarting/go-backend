package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"go-backend/internal/persistence/db/mocks"
	"go-backend/internal/persistence/db/model"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var crypto = model.Crypto{
	Name: "btc",
}

func TestCryptoHandler_Get(t *testing.T) {
	cryptoStoreMock := new(mocks.CryptoStoreMock)
	c := NewCryptoHandler(cryptoStoreMock)
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
				return cryptoStoreMock.On("GetByUuid", uuid).Return(model.Crypto{}, errors.New("db err"))
			},
			expectError: true,
		},
		{
			name: "Should return crypto",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", uuid.String())

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("GetByUuid", uuid).Return(crypto, nil)
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

			fn := c.Get()
			fn(tt.rec, tt.reqFn())

			if tt.expectError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				var result model.Crypto
				json.NewDecoder(tt.rec.Body).Decode(&result)

				assert.Equal(t, crypto, result)
				assert.Equal(t, "application/json", tt.rec.Header().Get("Content-Type"))
			}
			cryptoStoreMock.AssertExpectations(t)
		})
	}
}

func TestCryptoHandler_Create(t *testing.T) {
	cryptoStoreMock := new(mocks.CryptoStoreMock)
	c := NewCryptoHandler(cryptoStoreMock)

	tests := []struct {
		name         string
		rec          *httptest.ResponseRecorder
		arg          model.Crypto
		reqFn        func() *http.Request
		mockFn       func() *mock.Call
		expectResult bool
		expectError  bool
	}{
		{
			name: "Should create crypto",
			rec:  httptest.NewRecorder(),
			arg:  crypto,
			reqFn: func() *http.Request {
				r := httptest.NewRequest("POST", "/", nil)
				r.Form = url.Values{
					"name": {"btc"},
				}
				return r
			},
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("Create", mock.Anything).Return(nil)
			},
			expectResult: true,
			expectError:  false,
		},
		{
			name: "Should get error",
			rec:  httptest.NewRecorder(),
			arg:  model.Crypto{},
			reqFn: func() *http.Request {
				r := httptest.NewRequest("POST", "/", nil)
				r.Form = url.Values{
					"name": {"btc"},
				}
				return r
			},
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("Create", mock.Anything).Return(errors.New("failed to exec query"))
			},
			expectError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockCall *mock.Call
			if tt.mockFn != nil {
				mockCall = tt.mockFn()
			}

			fn := c.Create()
			fn(tt.rec, tt.reqFn())

			if tt.expectError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				assert.Equal(t, http.StatusOK, tt.rec.Code)
			}
			cryptoStoreMock.AssertExpectations(t)
			mockCall.Unset()
		})
	}
}

func TestCryptoHandler_Delete(t *testing.T) {
	cryptoStoreMock := new(mocks.CryptoStoreMock)
	c := NewCryptoHandler(cryptoStoreMock)

	uuid := uuid.New()
	tests := []struct {
		name         string
		rec          *httptest.ResponseRecorder
		arg          model.Crypto
		reqFn        func() *http.Request
		mockFn       func() *mock.Call
		expectResult bool
		expectError  bool
	}{
		{
			name: "Should delete crypto",
			rec:  httptest.NewRecorder(),
			arg:  crypto,
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", uuid.String())

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("Delete", uuid).Return(nil)
			},
			expectResult: true,
			expectError:  false,
		},
		{
			name: "Should get error when parsing uuid",
			rec:  httptest.NewRecorder(),
			arg:  model.Crypto{},
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
			name: "Should get error when deleting",
			rec:  httptest.NewRecorder(),
			arg:  model.Crypto{},
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", uuid.String())

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			mockFn: func() *mock.Call {
				return cryptoStoreMock.On("Delete", uuid).Return(errors.New("failed to exec query"))
			},
			expectError: true,
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

			fn := c.Delete()
			fn(tt.rec, tt.reqFn())

			if tt.expectError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				assert.Equal(t, http.StatusOK, tt.rec.Code)
			}
			cryptoStoreMock.AssertExpectations(t)
		})
	}
}
