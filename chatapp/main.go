package main

import (
	"log"
	"net/http"
	"real-time-chat/handlers"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	// Route to handle WebSocket connections
	http.HandleFunc("/ws", handlers.HandleConnections)
	// Start the message broadcasting loop in a separate goroutine
	go handlers.HandleMessages()
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
