package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)

type createLikeRequest struct {
	Id int64 `uri:"id" binding:"required"`
}
func (server *Server) createLike(ctx *gin.Context) {
	var req createLikeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {		
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	like, err := server.store.CreateLikeWithPost(ctx, db.CreateLikeWithPostParams{
		PostID: req.Id,
		Username: authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, like)
}

type deleteLikeRequest struct {
	Id int64 `uri:"id" binding:"required"`
}
func (server *Server) deleteLike(ctx *gin.Context) {
	var req deleteLikeRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	err := server.store.DeleteLikeWithPost(ctx, db.DeleteLikeWithPostParams{
		PostID: req.Id,
		Username: authPayload.Username,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"message": "success"})
}