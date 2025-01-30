package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Upgrader config to upgrade HTTP connections to WebSocket
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Connected clients map
var clients = make(map[*websocket.Conn]bool) // Connected clients
var broadcast = make(chan Message)           // Channel for broadcasting messages

// Message struct defines the structure of the messages exchanged
type Message struct {
	Username  string `json:"username"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	Typing    bool   `json:"typing"` // Indicates if the user is typing
}

// HandleConnections manages incoming WebSocket connections
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP to WebSocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer ws.Close() // Close WebSocket connection when the function exits

	clients[ws] = true // Add the new client

	// Listen for messages from this client
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Println("WebSocket read error:", err)
			delete(clients, ws) // Remove the client on error
			break
		}
		broadcast <- msg // Send the message to the broadcast channel
	}
}

// HandleMessages broadcasts messages to all connected clients
func HandleMessages() {
	for {
		// Get the next message from the broadcast channel
		msg := <-broadcast
		// Send the message to all clients
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Println("WebSocket write error:", err)
				client.Close()          // Close connection on error
				delete(clients, client) // Remove the client
			}
		}
	}
}
