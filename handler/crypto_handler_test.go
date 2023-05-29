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

var crypto = db.Crypto{}

func TestCryptoHandler_Get(t *testing.T) {
	cryptoStoreMock := new(mocks.CryptoStoreMock)
	c := &CryptoHandler{
		Store: cryptoStoreMock,
	}
	uuid := uuid.New()

	tests := []struct {
		name         string
		rec          *httptest.ResponseRecorder
		reqFn        func() *http.Request
		mockFn       func() *mock.Call
		expectResult bool
		exportError  bool
	}{
		{
			name: "Should return err if header is no uuid",
			rec:  httptest.NewRecorder(),
			reqFn: func() *http.Request {
				r := httptest.NewRequest("GET", "/", nil)

				rctx := chi.NewRouteContext()
				rctx.URLParams.Add("id", "malformed")

				r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
				return r
			},
			exportError: true,
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
				return cryptoStoreMock.On("GetByUuid", uuid).Return(db.Crypto{}, errors.New("db err"))
			},
			exportError: true,
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

			if tt.exportError {
				assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
				return
			}

			if tt.expectResult {
				result := db.Crypto{}
				json.NewDecoder(tt.rec.Body).Decode(result)

				assert.Equal(t, crypto, result)
				assert.Equal(t, "application/json", tt.rec.Header().Get("Content-Type"))
			}
			cryptoStoreMock.AssertExpectations(t)
		})
	}
}

// TODO WIP
//func TestCryptoHandler_Create(t *testing.T) {
//	cryptoStoreMock := new(mocks.CryptoStoreMock)
//	c := &CryptoHandler{
//		Store: cryptoStoreMock,
//	}
//	tests := []struct {
//		name         string
//		rec          *httptest.ResponseRecorder
//		reqFn        func() *http.Request
//		mockFn       func() *mock.Call
//		expectResult bool
//		exportError  bool
//	}{
//		{
//			name: "Should create crypto",
//			rec:  httptest.NewRecorder(),
//			reqFn: func() *http.Request {
//				r := httptest.NewRequest("POST", "/", nil)
//				r.URL.Query()
//				r.Form.Set("name", "btc")
//				return r
//			},
//			mockFn: func() *mock.Call {
//				return cryptoStoreMock.On("Create", db.Crypto{}).Return(nil)
//			},
//			expectResult: true,
//		},
//	}
//	for _, tt := range tests {
//		var mockCall *mock.Call
//		if tt.mockFn != nil {
//			mockCall = tt.mockFn()
//		}
//		defer func() {
//			if mockCall != nil {
//				mockCall.Unset()
//			}
//		}()
//
//		fn := c.Get()
//		fn(tt.rec, tt.reqFn())
//
//		if tt.exportError {
//			assert.Equal(t, http.StatusBadRequest, tt.rec.Code)
//			return
//		}
//
//		if tt.expectResult {
//			result := db.Crypto{}
//			json.NewDecoder(tt.rec.Body).Decode(result)
//
//			assert.Equal(t, crypto, result)
//			assert.Equal(t, "application/json", tt.rec.Header().Get("Content-Type"))
//		}
//		cryptoStoreMock.AssertExpectations(t)
//	}
//}
