package api

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zzooman/zapp-server/token"
)

// Client represents a single chatting client
type Client struct {
    ID     string
    Conn   *websocket.Conn
    RoomID string
}

// Room represents a chat room
type Room struct {
    ID      string
    Clients map[string]*Client
    Mutex   sync.Mutex
}
var rooms = make(map[string]*Room)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,	
}

func (server *Server) handleWebSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer conn.Close()
	roomID := ctx.Query("room_id")
	if roomID == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(errors.New("query parameter room_id is required")))
		return
	}
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	client := &Client{
		ID:     username,
		Conn:   conn,
		RoomID: roomID,
	}
	addClientToRoom(roomID, client)	
	defer removeClientFromRoom(roomID, username)
	for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            break
        }
        broadcastMessageToRoom(roomID, messageType, message)
    }
}

func broadcastMessageToRoom(roomID string, messageType int, message []byte) {
    room, exists := rooms[roomID]
    if exists {
        room.Mutex.Lock()
        defer room.Mutex.Unlock()
        for _, client := range room.Clients {
            err := client.Conn.WriteMessage(messageType, message)
            if err != nil {
                client.Conn.Close()
                delete(room.Clients, client.ID)
            }
        }
    }
}

func addClientToRoom(roomID string, client *Client) {
	room, ok := rooms[roomID]
	if !ok {
		room = &Room{
			ID:      roomID,
			Clients: make(map[string]*Client),
		}
		rooms[roomID] = room
	}	
	room.Mutex.Lock()
	room.Clients[client.ID] = client
	room.Mutex.Unlock()
}
func removeClientFromRoom(roomID string, clientID string) {
	room, ok := rooms[roomID]
	if !ok {
		return
	}
	room.Mutex.Lock()
	delete(room.Clients, clientID)
	room.Mutex.Unlock()
}