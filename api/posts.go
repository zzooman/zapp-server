package api

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type createPostRequest struct {
	Title   string   	`json:"title" binding:"required"`
	Content string   	`json:"content" binding:"required"`
	Medias  []*multipart.FileHeader 	`form:"medias"`
	Price   int64		`json:"price" binding:"required"`
	Stock   int64		`json:"stock" binding:"required"`
}
func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	// auth_payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)		
	fmt.Println(req.Medias)
	// post, err := server.store.CreatePost(ctx, db.CreatePostParams{
	// 	Author: auth_payload.Username,
	// 	Title:  req.Title,
	// 	Content: req.Content,
	// 	Medias: req.Medias,
	// 	Price:  req.Price,
	// 	Stock: req.Stock,		
	// })
}