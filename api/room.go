package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)



type enterChatRoomRequest struct {
	User_a string `json:"user_a" binding:"required"`
	User_b string `json:"user_b" binding:"required"`
}
func (server *Server) enterChatRoom(ctx *gin.Context) {
	var req enterChatRoomRequest
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
	if user_a.Username == user_b.Username {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
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

type getAllRoomsResponse struct {
	RoomID 			int64 	`json:"room_id"`
	Recipient 		string 	`json:"recipient"`
	LastMessage 	string 	`json:"last_message"`
	LastMessageAt 	string 	`json:"last_message_at"`
	UnreadCount 	int64 	`json:"unread_count"`	
}
func (server *Server) getAllRooms(ctx *gin.Context) {
	var response []getAllRoomsResponse
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	rooms, err := server.store.GetRoomsByUser(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	for i := range rooms {
		lastMessage, err := server.store.GetLastMessage(ctx, rooms[i].ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}
		unreadCount, err := server.store.UnreadMessageCount(ctx, username)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
			return
		}		
		var recipient string
		if rooms[i].UserA == username {
			recipient = rooms[i].UserB
		} else {
			recipient = rooms[i].UserA
		}
		response = append(response, getAllRoomsResponse{
					RoomID: rooms[i].ID,
					LastMessage: lastMessage.Message,
					LastMessageAt: lastMessage.CreatedAt.Time.String(),
					Recipient: recipient,
					UnreadCount: unreadCount,
		})		
	}	
	ctx.JSON(http.StatusOK, response)
}