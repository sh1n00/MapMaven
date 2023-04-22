package main

import (
	"backend/Handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", Handlers.HealthCheck)
	http.HandleFunc("/chat", Handlers.Chat)
	http.HandleFunc("/embeddings", Handlers.Embeddings)
	http.ListenAndServe("localhost:8080", nil)
}
