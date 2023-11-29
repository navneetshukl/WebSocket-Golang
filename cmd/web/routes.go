package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/navneetshukl/websocket/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndPoint))

	return mux
}
