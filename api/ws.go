package api

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zzooman/zapp-server/token"
)

// Client represents a single chatting client
type Client struct {
    Username     string
	RoomID int64
    Conn   *websocket.Conn    
}

// Room represents a chat room
type Room struct {
    ID      int64
    Clients map[string]*Client
    Mutex   sync.Mutex
}
var rooms = make(map[int64]*Room)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,	
}

type connectWebSocketRequest struct {
	RoomID int64 `json:"room_id" binding:"required"`
}
func (server *Server) connectWS(ctx *gin.Context) {
	var req connectWebSocketRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	defer conn.Close()

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	client := &Client{
		Username:     username,
		Conn:   conn,
		RoomID: req.RoomID,
	}
	enterRoom(req.RoomID, client)	
	defer exitRoom(req.RoomID, username)
	for {						
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            break
        }
		broadcastMessageToRoom(req.RoomID, messageType, message)   
    }
}

func broadcastMessageToRoom(roomID int64, messageType int, message []byte) {
    room, exists := rooms[roomID]
    if exists {
        room.Mutex.Lock()
        defer room.Mutex.Unlock()
        for _, client := range room.Clients {
            err := client.Conn.WriteMessage(messageType, message)
            if err != nil {
                client.Conn.Close()
                delete(room.Clients, client.Username)
            }
        }
    }
}

func enterRoom(roomID int64, client *Client) {
	room, ok := rooms[roomID]
	if !ok {
		room = &Room{
			ID:      roomID,
			Clients: make(map[string]*Client),
		}
		rooms[roomID] = room
	}	
	room.Mutex.Lock()
	room.Clients[client.Username] = client
	room.Mutex.Unlock()
}
func exitRoom(roomID int64, clientID string) {
	room, ok := rooms[roomID]
	if !ok {
		return
	}
	room.Mutex.Lock()
	delete(room.Clients, clientID)
	room.Mutex.Unlock()
}