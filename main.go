package main

import (
	"go-backend/api"
	_ "go-backend/queue"
	"net/http"
)

func main() {
	h := api.NewHandler()
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}
