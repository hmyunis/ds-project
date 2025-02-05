package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func testWebSocketServer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Simple echo response for testing
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(websocket.TextMessage, msg)
	}
}

func TestWebSocketConnection(t *testing.T) {
	// Start a test WebSocket server
	server := httptest.NewServer(http.HandlerFunc(testWebSocketServer))
	defer server.Close()

	wsURL := "ws" + server.URL[4:] // Convert http:// to ws://
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err, "WebSocket connection failed")

	defer conn.Close()

	// Send a test message
	message := "Hello WebSocket"
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	assert.NoError(t, err, "Failed to send message")

	// Read the response
	_, response, err := conn.ReadMessage()
	assert.NoError(t, err, "Failed to read message")
	assert.Equal(t, message, string(response), "Response did not match")
}
