package main

import (
	_ "go-backend/queue"
	"go-backend/route"
	"net/http"
)

func main() {
	h := route.NewHandler()
	_ = h.CreateAllRoutes()
	_ = http.ListenAndServe(":8888", h)
}
