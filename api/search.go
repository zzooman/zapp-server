package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/zzooman/zapp-server/db/sqlc"
)

type SearchPostsRequest struct {
	Query  string `form:"query" binding:"required"`
	Limit  int32  `form:"limit" binding:"required"`
	Page   int32  `form:"page" binding:"required"`
}
type SearchPostsResponse struct {
	Posts 	 []db.GetPostsWithAuthorByQueryRow 	`json:"posts"`
	Next  	 bool							  	`json:"next"`
	Query  	 string 							`json:"keyword"`
}
func (server *Server) searchPosts(ctx *gin.Context) {
	var req SearchPostsRequest
	var res SearchPostsResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}		

	result, err := server.store.SearchPostsTx(ctx, db.SearchPostsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	_, err = server.store.SearchPostsTx(ctx, db.SearchPostsParams{
		Query:  req.Query,
		Limit:  req.Limit,
		Offset: req.Page * req.Limit,
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
