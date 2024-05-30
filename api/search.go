package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

type SearchPostsRequest struct {
	Query  string `uri:"query" binding:"required"`
	Limit  int32  `uri:"limit" binding:"required"`
	Offset int32  `uri:"offset" binding:"required"`
}
type SearchPostsResult struct {
	Posts []db.GetPostsWithAuthorByQueryRow `json:"posts"`
}
func (server *Server) searchPosts(ctx *gin.Context) {
	var req SearchPostsRequest
	var res SearchPostsResult

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.store.SearchPostsTx(ctx, db.SearchPostsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res.Posts = result.Posts
	ctx.JSON(http.StatusOK, res)
}