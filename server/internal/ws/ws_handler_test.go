package ws

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHubForHandler struct {
	mock.Mock
	Rooms map[string]*Room
}

func NewMockHubForHandler() *MockHubForHandler {
	return &MockHubForHandler{
		Rooms: make(map[string]*Room),
	}
}

func (m *MockHubForHandler) Run() {
	m.Called()
	m.Rooms = make(map[string]*Room)
}

func (m *MockHubForHandler) AddRoom(room *Room) {
	m.Rooms[room.ID] = room
}

func (m *MockHubForHandler) GetRooms() map[string]*Room {
	return m.Rooms
}

func TestGetRooms(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockHub := NewMockHubForHandler()
	handler := NewHandler(&Hub{Rooms: mockHub.Rooms})
	mockHub.AddRoom(&Room{ID: "testRoom1", Name: "Test Room 1"})
	mockHub.AddRoom(&Room{ID: "testRoom2", Name: "Test Room 2"})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req, _ := http.NewRequest("GET", "/ws/getRooms", nil)
	c.Request = req

	handler.GetRooms(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":"testRoom1"`)
	assert.Contains(t, w.Body.String(), `"id":"testRoom2"`)
}

func TestGetClients(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockHub := NewMockHubForHandler()
	handler := NewHandler(&Hub{Rooms: mockHub.Rooms})

	roomID := "testRoom"
	mockHub.AddRoom(&Room{
		ID:   roomID,
		Name: "Test Room",
		Clients: map[string]*Client{
			"client1": {ID: "client1", Username: "User1"},
			"client2": {ID: "client2", Username: "User2"},
		},
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "roomId", Value: roomID}}
	req, _ := http.NewRequest("GET", "/ws/getClients/"+roomID, nil)
	c.Request = req

	handler.GetClients(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"id":"client1"`)
	assert.Contains(t, w.Body.String(), `"id":"client2"`)
}

func TestGetMessagesInRoom(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockHub := NewMockHubForHandler()
	handler := NewHandler(&Hub{Rooms: mockHub.Rooms})

	roomID := "testRoom"
	mockHub.AddRoom(&Room{
		ID:   roomID,
		Name: "Test Room",
		Messages: []*Message{
			{Content: "Message 1", Username: "User1"},
			{Content: "Message 2", Username: "User2"},
		},
	})

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "roomId", Value: roomID}}
	req, _ := http.NewRequest("GET", "/ws/getMessages/"+roomID, nil)
	c.Request = req

	handler.GetMessagesInRoom(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"content":"Message 1"`)
	assert.Contains(t, w.Body.String(), `"content":"Message 2"`)
}
