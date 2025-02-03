package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(h *Hub) *Handler {
	return &Handler{
		hub: h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetMessagesInRoom(c *gin.Context) {
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; ok {
		c.JSON(http.StatusOK, h.hub.Rooms[roomId].Messages)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var req CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRoom := &Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*Client),
	}

	h.hub.Rooms[req.ID] = newRoom

	if err := SaveRoom(newRoom); err != nil { //Changed from db.SaveRoom
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save room to DB"})
		return
	}

	c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Param("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	if _, ok := h.hub.Rooms[roomID]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room does not exist"})
		return
	}

	cl := &Client{
		Conn:     conn,
		Message:  make(chan *Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
		return
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}

func (h *Handler) DeleteRoom(c *gin.Context) {
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	delete(h.hub.Rooms, roomId)

	if err := DeleteRoom(roomId); err != nil { //Changed from db.DeleteRoom
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete room from DB"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}
