package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/goioc/di"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-backend/db"
	"go-backend/goroutines"
	"go-backend/handler"
	"go-backend/route"
	"go-backend/service"
	"net/http"
	"reflect"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store, err := db.NewStore("postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err)
	}

	registerDependencies(store)

	if err := di.InitializeContainer(); err != nil {
		log.Fatal().Err(err)
		return
	}

	goroutines.StartDispatcher(1)

	// Create handler and listen+serve for requests in a blocking manner
	h := di.GetInstance("handler").(*route.Handler)
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}

func registerDependencies(store *db.Store) {
	_, _ = di.RegisterBeanInstance("store", store)
	_, _ = di.RegisterBeanInstance("walletStore", store.WalletStore)
	_, _ = di.RegisterBeanInstance("cryptoStore", store.CryptoStore)
	_, _ = di.RegisterBeanInstance("walletCryptoStore", store.WalletCryptoStore)

	_, _ = di.RegisterBean("walletHandler", reflect.TypeOf((*handler.WalletHandler)(nil)))
	_, _ = di.RegisterBean("cryptoHandler", reflect.TypeOf((*handler.CryptoHandler)(nil)))
	_, _ = di.RegisterBean("withdrawalHandler", reflect.TypeOf((*handler.WithdrawalHandler)(nil)))

	_, _ = di.RegisterBean("withdrawalService", reflect.TypeOf((*service.WithdrawalService)(nil)))

	_, _ = di.RegisterBeanFactory("handler", di.Singleton, func(ctx context.Context) (interface{}, error) {
		return &route.Handler{
			Mux: chi.NewMux(),
		}, nil
	})
}
