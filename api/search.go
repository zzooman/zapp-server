package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

type SearchPostsRequest struct {
	Query  string `form:"query" binding:"required"`
	Limit  int32  `form:"limit" binding:"required"`
	Offset int32  `form:"offset" binding:"required"`
}

type SearchPostsResult struct {
	Posts []db.GetPostsWithAuthorByQueryRow `json:"posts"`
}

func (server *Server) searchPosts(ctx *gin.Context) {
	var req SearchPostsRequest
	var res SearchPostsResult

	// Set default values for Limit and Offset
	limitStr := ctx.DefaultQuery("limit", "10")
	offsetStr := ctx.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req.Query = ctx.Query("query")
	req.Limit = int32(limit)
	req.Offset = int32(offset)

	if req.Query == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("query is required")))
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
