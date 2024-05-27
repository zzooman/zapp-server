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
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Price   string   `json:"price" binding:"required"`
	Stock   int64    `json:"stock" binding:"required"`
	Medias  []string `json:"medias" binding:"required"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req createPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(AUTH_TOKEN).(*token.Payload)
	price, _ := strconv.ParseInt(req.Price, 10, 64)
	post, err := server.store.CreatePost(ctx, db.CreatePostParams{
		Author:    authPayload.Username,
		Title:     req.Title,
		Content:   req.Content,
		Price:     price,
		Stock:     req.Stock,
		Medias:     req.Medias,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, post)
}

type getPostsRequest struct {
	Limit  int32 `form:"limit"`
	Offset int32 `form:"offset"`
}

type Author struct {
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Phone    pgtype.Text `json:"phone"`
	Profile  pgtype.Text `json:"profile"`
}

type PostResponse struct {
	ID        int64               `json:"id"`
	Title     string              `json:"title"`
	Content   string              `json:"content"`
	Medias    []string            `json:"medias"`
	Price     int64               `json:"price"`
	Stock     int64               `json:"stock"`
	Views     pgtype.Int8         `json:"views"`
	CreatedAt pgtype.Timestamptz  `json:"created_at"`
	Author    Author              `json:"author"`
	IsLiked   bool                `json:"isLiked"`
}

func (server *Server) getPosts(ctx *gin.Context) {
	var req getPostsRequest
	var res []PostResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 게시글 & 작성자 정보 조회
	postsWithAuthor, err := server.store.GetPostsWithAuthor(ctx, db.GetPostsWithAuthorParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 채널 생성
	ch := make(chan struct {
		db.GetPostsWithAuthorRow
		IsLiked bool
	}, len(postsWithAuthor))

	// 비동기 처리
	for _, post := range postsWithAuthor {
		go func(post db.GetPostsWithAuthorRow) {			
			_, err := server.store.GetLikeWithPost(ctx, db.GetLikeWithPostParams{
				PostID:   post.ID,
				Username: post.Author,
			})	
			ch <- struct {
				db.GetPostsWithAuthorRow
				IsLiked bool
			}{post, err == nil}
		}(post)
	}

	// 결과 수집
	for range postsWithAuthor {
		result := <-ch
		res = append(res, PostResponse{
			ID:        result.ID,
			Title:     result.Title,
			Content:   result.Content,
			Medias:     result.Medias,
			Price:     result.Price,
			Stock:     result.Stock,
			Views:     result.Views,
			CreatedAt: result.CreatedAt,
			Author: Author{
				Username: result.Author,
				Email:    result.Email,
				Phone:    result.Phone,
				Profile:  result.Profile,
			},
			IsLiked: result.IsLiked,
		})
	}

	ctx.JSON(http.StatusOK, res)
}
