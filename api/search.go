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
	Page int32  `form:"page" binding:"required"`
}
type SearchPostsResponse struct {
	Posts []db.GetPostsWithAuthorByQueryRow `json:"posts"`
	Next  bool							  	`json:"next"`
}
func (server *Server) searchPosts(ctx *gin.Context) {
	var req SearchPostsRequest
	var res SearchPostsResponse

	// Set default values for Limit and Page
	limitStr := ctx.DefaultQuery("limit", "10")
	pageStr := ctx.DefaultQuery("page", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req.Query = ctx.Query("query")
	req.Limit = int32(limit)
	req.Page = int32(page)

	if req.Query == "" {
		ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("query is required")))
		return
	}

	result, err := server.store.SearchPostsTx(ctx, db.SearchPostsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: req.Page * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.SearchPostsTx(ctx, db.SearchPostsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: (req.Page + 1) * req.Limit,
	})
	if err != nil { res.Next = false } else { res.Next = true }

	res.Posts = result.Posts
	ctx.JSON(http.StatusOK, res)
}

func (server *Server) hotSearchTexts(ctx *gin.Context) {
	searchTexts, err := server.store.HotSearchTexts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, searchTexts)
}
