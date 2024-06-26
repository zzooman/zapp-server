package api

import (
	"fmt"
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
// 포스트 조회
type getPostRequest struct {
	Id string `uri:"id" binding:"required"`	
}
func (server *Server) getPost(ctx *gin.Context) {
	var req getPostRequest	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	postWithAuthor, err := server.store.GetPostWithAuthor(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	_, err = server.store.GetLikeWithPost(ctx, db.GetLikeWithPostParams{
		PostID:   postWithAuthor.ID,
		Username: postWithAuthor.Author,
	})
	ctx.JSON(http.StatusOK, PostResponse{
		ID:        	postWithAuthor.ID,
		Title:     	postWithAuthor.Title,
		Content:   	postWithAuthor.Content,
		Medias:    	postWithAuthor.Medias,
		Price:     	postWithAuthor.Price,
		Stock:     	postWithAuthor.Stock,
		Views:     	postWithAuthor.Views,
		CreatedAt: 	postWithAuthor.CreatedAt,
		Author: Author{
			Username: postWithAuthor.Author,
			Email:    postWithAuthor.Email,
			Phone:    postWithAuthor.Phone,
			Profile:  postWithAuthor.Profile,			
		},
		IsLiked: err == nil,
	})
}

type getPostsRequest struct {
	Limit  	int32 	`form:"limit"`
	Page 	int32 	`form:"page"`
}
type getPostsResponse struct {
	Posts []PostResponse 	`json:"posts"`
	Next  bool		   		`json:"next"`
}
// 포스트 목록 조회
func (server *Server) getPosts(ctx *gin.Context) {
	var req getPostsRequest
	var res getPostsResponse

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 게시글 & 작성자 정보 조회
	postsWithAuthor, err := server.store.GetPostsWithAuthor(ctx, db.GetPostsWithAuthorParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 다음 페이지 존재 여부 확인
	nextPosts, err := server.store.GetPosts(ctx, db.GetPostsParams{
		Limit:  req.Limit,
		Offset: req.Page * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if len(nextPosts) == 0 {
		res.Next = false
	} else {
		res.Next = true	
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
		res.Posts = append(res.Posts, PostResponse{
			ID:        	result.ID,
			Title:     	result.Title,
			Content:   	result.Content,
			Medias:     result.Medias,
			Price:     	result.Price,
			Stock:     	result.Stock,
			Views:     	result.Views,
			CreatedAt: 	result.CreatedAt,
			Author: Author{
				Username: result.Author,
				Email:    result.Email,
				Phone:    result.Phone,
				Profile:  result.Profile,
			},
			IsLiked: result.IsLiked,
		})
	}

	posts := res.Posts
	// 최신순 정렬
	for range posts {
		for i := 0; i < len(posts)-1; i++ {
			if posts[i].CreatedAt.Time.Before(posts[i+1].CreatedAt.Time) {				
				posts[i], posts[i+1] = posts[i+1], posts[i]
			}
		}
	}	
	ctx.JSON(http.StatusOK, res)	
}

// TODO : 포스트 수정
// TODO : 포스트 삭제

func (server *Server) getPostsILiked(ctx *gin.Context) {
	var req getPostsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	posts, err := server.store.GetPostsWithAuthorThatILiked(ctx, db.GetPostsWithAuthorThatILikedParams{
		Username: 	username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}


func (server *Server) getPostsISold(ctx *gin.Context) {
	var req getPostsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	posts, err := server.store.GetPostsWithAuthorThatISold(ctx, db.GetPostsWithAuthorThatISoldParams{
		Seller: 	username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (server *Server) getPostsIBought(ctx *gin.Context) {
	var req getPostsRequest
	fmt.Println("start")
	if err := ctx.ShouldBindQuery(&req); err != nil {
		fmt.Println("error 1", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	fmt.Println("req", req)
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	fmt.Println("username", username)
	posts, err := server.store.GetPostsWithAuthorThatIBought(ctx, db.GetPostsWithAuthorThatIBoughtParams{		
		Buyer: 		username,
		Limit:    	req.Limit,
		Offset:  	(req.Page - 1) * req.Limit,
	})
	if err != nil {
		fmt.Println("error 2", err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	fmt.Println("posts", posts)
	ctx.JSON(http.StatusOK, posts)
}