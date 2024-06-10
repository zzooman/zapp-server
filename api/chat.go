package api

import (
	"errors"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	db "github.com/zzooman/zapp-server/db/sqlc"
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
	enterRoom(roomID, client)	
	defer exitRoom(roomID, username)
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

func enterRoom(roomID string, client *Client) {
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
func exitRoom(roomID string, clientID string) {
	room, ok := rooms[roomID]
	if !ok {
		return
	}
	room.Mutex.Lock()
	delete(room.Clients, clientID)
	room.Mutex.Unlock()
}

type createRoomRequest struct {
	User_a string `json:"user_a" binding:"required"`
	User_b string `json:"user_b" binding:"required"`
}
func (server *Server) createRoom(ctx *gin.Context) {
	var req createRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	user_a, err := server.store.GetUser(ctx, req.User_a)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	user_b, err := server.store.GetUser(ctx, req.User_b)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	existingRoom, err := server.store.CheckRoom(ctx, db.CheckRoomParams{
		UserA: user_a.Username,
		UserB: user_b.Username,
	})	
	if err == nil {
		ctx.JSON(http.StatusOK, existingRoom)
		return
	}
	room, err := server.store.CreateRoom(ctx, db.CreateRoomParams{
		UserA: user_a.Username,
		UserB: user_b.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, room)
}