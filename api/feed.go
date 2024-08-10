package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zzooman/zapp-server/db/sqlc"
	"github.com/zzooman/zapp-server/token"
)

type Author struct {	
	Username 	string      `json:"username"`
	Email    	string      `json:"email"`
	Phone    	pgtype.Text `json:"phone"`
	Profile  	pgtype.Text `json:"profile"`
}
type FeedResponse struct {
	ID          int64               `json:"id"`	
	Content     string              `json:"content"`
	Medias      []string            `json:"medias"`	
	Views       pgtype.Int8         `json:"views"`
	CreatedAt   pgtype.Timestamptz  `json:"created_at"`
	Author      Author              `json:"author"`
	IsLiked     bool                `json:"isLiked"`
	Comments   	int32				`json:"comments"`
}
type GetFeedsRequest struct {
	Limit int32 `form:"limit" binding:"required"`
	Page  int32 `form:"page" binding:"required"`
}
type GetFeedsResponse struct {
	Next bool				`json:"next"`
	Feeds []FeedResponse    `json:"feeds"`
}
func (server *Server) getFeeds(ctx *gin.Context) {
	var req GetFeedsRequest
	var res GetFeedsResponse	

	if err := ctx.BindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	feedsWithAuthor, err := server.store.GetFeedsWithAuthor(ctx, db.GetFeedsWithAuthorParams{
		Limit: req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
		
	nextFeeds, err := server.store.GetFeeds(ctx, db.GetFeedsParams{
		Limit: req.Limit,
		Offset: req.Limit * req.Page,
	})
	if len(nextFeeds) == 0 {
		res.Next = false
	} else {
		res.Next = true
	}

	ch := make(chan struct {
		db.GetFeedsWithAuthorRow
		IsLiked 	bool
		Comments 	int32
	}, len(feedsWithAuthor))
	
	var username string = ""
	auth, ok := ctx.Get(AUTH_TOKEN)
	if ok {
		username = auth.(*token.Payload).Username
	}
	
	for _, feed := range feedsWithAuthor {
		go func(feed db.GetFeedsWithAuthorRow) {
			_, err = server.store.GetLikeWithFeed(ctx, db.GetLikeWithFeedParams{
				Username: username,
				FeedID: feed.ID,
			})			
			count, err := server.store.GetCountOfComments(ctx, feed.ID)
			ch <- struct{db.GetFeedsWithAuthorRow; IsLiked bool; Comments int32;}{
				feed,
				err == nil,
				int32(count),
			}
		}(feed)
	} 

	for range feedsWithAuthor {
		result := <-ch
		res.Feeds = append(res.Feeds, FeedResponse{
			Author: Author{
				Username: result.Author,
				Phone: result.Phone,
				Profile: result.Profile,
				Email: result.Email,
			},
			ID: result.ID,
			Content: result.Content,
			Medias: result.Medias,
			Views: result.Views,
			CreatedAt: result.CreatedAt,
			IsLiked: result.IsLiked,
			Comments: result.Comments,
		})
	}

	feeds := res.Feeds
	// 최신순 버블 정렬
	for range feeds {
		for i := 0; i < len(feeds)-1; i++ {
			if feeds[i].CreatedAt.Time.Before(feeds[i+1].CreatedAt.Time) {				
				feeds[i], feeds[i+1] = feeds[i+1], feeds[i]
			}
		}
	}	
	ctx.JSON(http.StatusOK, res)	
}


type GetFeedDetailRequest struct {
	Id int `uri:"id" binding:"required"`
}
type Comment struct {
	ID              int64              `json:"id"`
    FeedID          int64              `json:"feed_id"`
    ParentCommentID pgtype.Int8        `json:"parent_comment_id"`
    Commentor       Author             `json:"commentor"`
    CommentText     string             `json:"comment_text"`
    CreatedAt       pgtype.Timestamptz `json:"created_at"`
}
type GetFeedDetailResponse struct {
	ID          int64               `json:"id"`	
	Content     string              `json:"content"`
	Medias      []string            `json:"medias"`	
	Views       pgtype.Int8         `json:"views"`
	CreatedAt   pgtype.Timestamptz  `json:"created_at"`
	Author      Author              `json:"author"`
	IsLiked     bool                `json:"isLiked"`
	Comments   	[]Comment			`json:"comments"`
}
func (server *Server) getFeed(ctx *gin.Context) {
	var req GetFeedDetailRequest
	var res GetFeedDetailResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	
	feedWithAuthor, err := server.store.GetFeedWithAuthor(ctx, int64(req.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.store.ViewFeed(ctx, feedWithAuthor.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	username := ""
	auth, ok := ctx.Get(AUTH_TOKEN)
	if ok {
		username = auth.(*token.Payload).Username
	}

	_, err = server.store.GetLikeWithFeed(ctx, db.GetLikeWithFeedParams{
		Username: username,
		FeedID: feedWithAuthor.ID,
	})

	res = GetFeedDetailResponse{
		ID: feedWithAuthor.ID,
		Content: feedWithAuthor.Content,
		Medias: feedWithAuthor.Medias,
		Views: feedWithAuthor.Views,
		CreatedAt: feedWithAuthor.CreatedAt,
		Author: Author{
			Username: feedWithAuthor.Author,
			Email: feedWithAuthor.Email,
			Phone: feedWithAuthor.Phone,
			Profile: feedWithAuthor.Profile,
		},
		IsLiked: err == nil,		
	}

	comments, err := server.store.GetComments(ctx, db.GetCommentsParams{
		FeedID: feedWithAuthor.ID,
		Limit: 10,
		Offset: 0,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}	

	if len(comments) == 0 {
		res.Comments = []Comment{}
	}

	ch := make(chan Comment, len(comments))

	for i := 0; i < len(comments); i++ {
		go func(comment db.Comment) {
			commentor, _ := server.store.GetUser(ctx, comment.Commentor)			
			ch <- Comment{
				ID: comment.ID,
				FeedID: comment.FeedID,              
				ParentCommentID: comment.ParentCommentID,
				Commentor: Author{
					Username: commentor.Username,       
    				Email: commentor.Email,          
    				Phone: commentor.Phone,    
    				Profile: commentor.Profile,  
				},             
				CommentText: comment.CommentText,             
				CreatedAt: comment.CreatedAt,
			}
			
		}(comments[i])
	}

	for range comments {
		result := <-ch
		res.Comments = append(res.Comments, result)
	}

	
	// 최신순 버블 정렬
	for range res.Comments {
		for i := 0; i < len(res.Comments)-1; i++ {
			if res.Comments[i].CreatedAt.Time.Before(res.Comments[i+1].CreatedAt.Time) {				
				res.Comments[i], res.Comments[i+1] = res.Comments[i+1], res.Comments[i]
			}
		}
	}	

	ctx.JSON(http.StatusOK, res)
}


type createFeedRequest struct {	
	Content string   `json:"content" binding:"required"`	
	Medias  []string `json:"medias"`
}

func (server *Server) createFeed(ctx *gin.Context) {
	var req createFeedRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	
	feed, err := server.store.CreateFeed(ctx, db.CreateFeedParams{
		Author: username,
		Content: req.Content,
		Medias: req.Medias,
		CreatedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, feed)
}