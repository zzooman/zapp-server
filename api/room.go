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