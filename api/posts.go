package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzooman/zapp-server/token"
)

type createPostRequest struct {
	Title   string   	`json:"title" binding:"required"`
	Content string   	`json:"content" binding:"required"`
	Medias  []string 	`json:"medias"`
	Price   int64		`json:"price"`
}
func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	auth_payload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	
	
}