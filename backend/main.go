package main

import (
	"backend/Handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health-check", Handlers.HealthCheck)
	http.HandleFunc("/chat", Handlers.Chat)
	http.HandleFunc("/text-to-audio", Handlers.TextToAudio)
	log.Println("Starting Server")
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatalln(err)
	}
}
