package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)



type enterChatRoomRequest struct {
	Host 		string 	`json:"host" binding:"required"`
	Guest 		string 	`json:"guest" binding:"required"`
	ProductId 	int 	`json:"product_id"`
}
func (server *Server) enterChatRoom(ctx *gin.Context) {
	var req enterChatRoomRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {		
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	host, err := server.store.GetUser(ctx, req.Host)
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	guest, err := server.store.GetUser(ctx, req.Guest)
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if host.Username == guest.Username {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	existingRoom, err := server.store.CheckRoom(ctx, db.CheckRoomParams{
		Host: host.Username,
		Guest: guest.Username,		
	})	
	if err == nil {		
		ctx.JSON(http.StatusOK, existingRoom)			
		return
	}
	
	room, err := server.store.CreateRoom(ctx, db.CreateRoomParams{
		Host: host.Username,
		Guest: guest.Username,
		ProductID: pgtype.Int8{Int64: int64(req.ProductId), Valid: false},		
	})
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}	
	ctx.JSON(http.StatusOK, room)
}

type getMessagesRequest struct {
	RoomID int64 `uri:"room_id" binding:"required"`
}
func (server *Server) getMessages(ctx *gin.Context) {
	var req getMessagesRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	messages, err := server.store.GetMessagesByRoom(ctx, req.RoomID)
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	for i := range messages {
		message := &messages[i]
		if message.Sender != username && !message.ReadAt.Valid {
			err = server.store.ReadMessage(ctx, message.ID)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errorResponse(err))
				return
			}
		}
	}	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, messages)
}

func (server *Server) createMessage(ctx *gin.Context, roomID int64, message string, sender string) {
	room, err := server.store.GetRoom(ctx, roomID)
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	newMessage, err := server.store.CreateMessage(ctx, db.CreateMessageParams{
		Sender:  sender,
		RoomID:  room.ID,
		Message: message,
	})
	if err != nil {		
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}	
	ctx.JSON(http.StatusOK, newMessage)
}


type RoomType string
const (
	All	  RoomType = "all"
	Chat  RoomType = "chat"
	Buy   RoomType = "buy"
	Sell  RoomType = "sell"
)

// GetRoomsRequest struct with custom RoomType
type GetRoomsRequest struct {
	Type RoomType `uri:"type" binding:"required,roomType"`
}
type RoomResponse struct {
	RoomID 			int64 	`json:"room_id"`
	Recipient 		string 	`json:"recipient"`
	LastMessage 	string 	`json:"last_message"`
	LastMessageAt 	string 	`json:"last_message_at"`
	UnreadCount 	int64 	`json:"unread_count"`	
}
func (server *Server) getRooms(ctx *gin.Context) {
	var res []RoomResponse
	var req GetRoomsRequest
	if err := ctx.ShouldBindUri(&res); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	

	if(req.Type == "all") {
		getAllRooms(ctx, server.store)
		return
	}
	if(req.Type == "chat") {
		getChatsRooms(ctx, server.store)
	}
	if(req.Type == "buy") {
		getBuyRooms(ctx, server.store)
	}
	if(req.Type == "sell") {
		getSellRooms(ctx, server.store)
	}
}

func getAllRooms(ctx *gin.Context, store db.Store) {
	var res []RoomResponse
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	rooms, err := store.GetRoomsByUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for i := range rooms {
		lastMessage, err := store.GetLastMessage(ctx, rooms[i].ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		unreadCount, err := store.UnreadMessageCount(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}		
		var recipient string
		if rooms[i].Host == username {
			recipient = rooms[i].Guest
		} else {
			recipient = rooms[i].Host
		}
		res = append(res, RoomResponse{
			RoomID: rooms[i].ID,
			LastMessage: lastMessage.Message,
			LastMessageAt: lastMessage.CreatedAt.Time.String(),
			Recipient: recipient,
			UnreadCount: unreadCount,
		})		
	}	
	ctx.JSON(http.StatusOK, res)
}


func getChatsRooms(ctx *gin.Context, store db.Store) {
	var res []RoomResponse
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username

	rooms, err := store.GetChatsRoomByUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	for i := range rooms {
		lastMessage, err := store.GetLastMessage(ctx, rooms[i].ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		unreadCount, err := store.UnreadMessageCount(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}		
		var recipient string
		if rooms[i].Host == username {
			recipient = rooms[i].Guest
		} else {
			recipient = rooms[i].Host
		}
		res = append(res, RoomResponse{
			RoomID: rooms[i].ID,
			LastMessage: lastMessage.Message,
			LastMessageAt: lastMessage.CreatedAt.Time.String(),
			Recipient: recipient,
			UnreadCount: unreadCount,
		})		
	}	
	ctx.JSON(http.StatusOK, res)
}

func getBuyRooms(ctx *gin.Context, store db.Store) {
	var res []RoomResponse
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username

	rooms, err := store.GetBuyRoomByUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	for i := range rooms {
		lastMessage, err := store.GetLastMessage(ctx, rooms[i].ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		unreadCount, err := store.UnreadMessageCount(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}		
		var recipient string
		if rooms[i].Host == username {
			recipient = rooms[i].Guest
		} else {
			recipient = rooms[i].Host
		}
		res = append(res, RoomResponse{
			RoomID: rooms[i].ID,
			LastMessage: lastMessage.Message,
			LastMessageAt: lastMessage.CreatedAt.Time.String(),
			Recipient: recipient,
			UnreadCount: unreadCount,
		})		
	}	
	ctx.JSON(http.StatusOK, res)
}

func getSellRooms(ctx *gin.Context, store db.Store) {
	var res []RoomResponse
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username

	rooms, err := store.GetSellRoomByUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	for i := range rooms {
		lastMessage, err := store.GetLastMessage(ctx, rooms[i].ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		unreadCount, err := store.UnreadMessageCount(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}		
		var recipient string
		if rooms[i].Host == username {
			recipient = rooms[i].Guest
		} else {
			recipient = rooms[i].Host
		}
		res = append(res, RoomResponse{
			RoomID: rooms[i].ID,
			LastMessage: lastMessage.Message,
			LastMessageAt: lastMessage.CreatedAt.Time.String(),
			Recipient: recipient,
			UnreadCount: unreadCount,
		})		
	}	
	ctx.JSON(http.StatusOK, res)
}