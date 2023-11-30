package main

import (
	"log"
	"net/http"

	"github.com/navneetshukl/websocket/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel Listener")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 8080")
	_=http.ListenAndServe(":8080",mux)
}
