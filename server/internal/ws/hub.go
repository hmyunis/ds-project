package ws

import (
	"log"
)

type Room struct {
	ID       string             `json:"id"`
	Name     string             `json:"name"`
	Clients  map[string]*Client `json:"clients"`
	Messages []*Message         `json:"messages"`
}

type Hub struct {
	Rooms      map[string]*Room
	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	rooms, err := GetRooms() //Changed from db.GetRooms
	if err != nil {
		log.Println("Failed to get rooms from db", err)
		rooms = make(map[string]*Room) // Initialize with empty map if there is an error
	}

	return &Hub{
		Rooms:      rooms,
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case cl := <-h.Register:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				r := h.Rooms[cl.RoomID]

				if _, ok := r.Clients[cl.ID]; !ok {
					r.Clients[cl.ID] = cl
					SaveRoom(r) //Changed from db.SaveRoom
				}
			}
		case cl := <-h.Unregister:
			if _, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := h.Rooms[cl.RoomID].Clients[cl.ID]; ok {
					if len(h.Rooms[cl.RoomID].Clients) != 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}
					delete(h.Rooms[cl.RoomID].Clients, cl.ID)
					close(cl.Message)
					err := SaveRoom(h.Rooms[cl.RoomID]) //Changed from db.SaveRoom
					if err != nil {
						log.Println("Failed to save room to db after user unregister:", err)
					}
				}
			}

		case m := <-h.Broadcast:
			if _, ok := h.Rooms[m.RoomID]; ok {
				h.Rooms[m.RoomID].Messages = append(h.Rooms[m.RoomID].Messages, m)
				for _, cl := range h.Rooms[m.RoomID].Clients {
					cl.Message <- m
				}
				SaveRoom(h.Rooms[m.RoomID]) //Changed from db.SaveRoom
			}
		}
	}
}
