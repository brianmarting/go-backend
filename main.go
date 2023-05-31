package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go-backend/db"
	"go-backend/route"
	"go-backend/worker"
	"net/http"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	store, err := db.NewStore()
	if err != nil {
		log.Fatal().Err(err)
	}

	worker.StartDispatcher(1)

	h := route.NewHandler(store)
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)

	// not being called?
	worker.StopWorkers()
}
