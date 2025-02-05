package tests

import (
	"net/http/httptest"
	"testing"

	"server/internal/ws"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type WebSocketTestSuite struct {
	suite.Suite
	server *httptest.Server
	url    string
}

func (suite *WebSocketTestSuite) SetupSuite() {
	// Initialize Gin server & WebSocket routes
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	hub := ws.NewHub()
	go hub.Run()

	handler := ws.NewHandler(hub)
	router.POST("/room", handler.CreateRoom)
	router.GET("/ws/:roomId", handler.JoinRoom)

	suite.server = httptest.NewServer(router)
	suite.url = "ws" + suite.server.URL[4:] // Change HTTP to WS
}

func (suite *WebSocketTestSuite) TearDownSuite() {
	suite.server.Close()
}

func (suite *WebSocketTestSuite) TestCreateRoom() {
	client := httptest.NewRequest("POST", "/room", nil)
	w := httptest.NewRecorder()
	suite.server.Config.Handler.ServeHTTP(w, client)

	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *WebSocketTestSuite) TestJoinRoom() {
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial(suite.url+"/ws/room1", nil)
	assert.NoError(suite.T(), err)
	defer conn.Close()
}

func TestWebSocketSuite(t *testing.T) {
	suite.Run(t, new(WebSocketTestSuite))
}
