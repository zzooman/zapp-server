package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)

type createPostRequest struct {
	Title   string   	`json:"title" binding:"required"`
	Content string   	`json:"content" binding:"required"`	
	Price   string		`json:"price" binding:"required"`
	Stock   int64		`json:"stock" binding:"required"`
	Medias  []string	`json:"medias" binding:"required"`
}
func (server *Server) createPost(ctx *gin.Context) {	
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {		
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	auth_payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)			
	price, _ := strconv.ParseInt(req.Price, 10, 64)
	post, err := server.store.CreatePost(ctx, db.CreatePostParams{
		Author:    	auth_payload.Username,
		Title:     	req.Title,
		Content:   	req.Content,		
		Price:     	price,
		Stock:     	req.Stock,
		Media:    	req.Medias,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},				
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}


