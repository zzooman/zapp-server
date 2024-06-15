package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
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