package api

import (
	"net/http"

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
	ID        int64               `json:"id"`	
	Content   string              `json:"content"`
	Medias    []string            `json:"medias"`	
	Views     pgtype.Int8         `json:"views"`
	CreatedAt pgtype.Timestamptz  `json:"created_at"`
	Author    Author              `json:"author"`
	IsLiked   bool                `json:"isLiked"`
}
type GetFeedsRequest struct {
	Limit 	int32 `json:"limit"`
	Page 	int32 `json:"page"`
}
type GetFeedsResponse struct {
	Next bool				`json:"next"`
	feeds []FeedResponse    `json:"feeds"`
}
func (server *Server) getFeeds(ctx *gin.Context) {
	var req GetFeedsRequest
	var res GetFeedsResponse
	
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return 
	}

	feedsWithAuthor, err := server.store.GetFeedsWithAuthor(ctx, db.GetFeedsWithAuthorParams{
		Limit: req.Limit,
		Offset: req.Limit - 1 * req.Page,
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
		IsLiked bool
	}, len(feedsWithAuthor))
	
	username := ctx.MustGet(AUTH_TOKEN).(*token.Payload).Username
	for _, feed := range feedsWithAuthor {
		go func(feed db.GetFeedsWithAuthorRow) {
			_, err = server.store.GetLikeWithFeed(ctx, db.GetLikeWithFeedParams{
				Username: username,
				FeedID: feed.ID,
			})
			ch <- struct{db.GetFeedsWithAuthorRow; IsLiked bool}{
				feed,
				err == nil,
			}
		}(feed)
	} 

	for range feedsWithAuthor {
		result := <-ch
		res.feeds = append(res.feeds, FeedResponse{
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
		})
	}

	feeds := res.feeds
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


type GetFeedRequest struct {
	Id int `json:"id"`	
}
func (server *Server) getFeed(ctx *gin.Context) {
	var req GetFeedRequest
	var res FeedResponse

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	feedWithAuthor, err := server.store.GetFeedWithAuthor(ctx, int64(req.Id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	_, err = server.store.GetLikeWithFeed(ctx, db.GetLikeWithFeedParams{
		Username: feedWithAuthor.Author,
		FeedID: feedWithAuthor.ID,
	})
	
	res = FeedResponse{
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

	ctx.JSON(http.StatusOK, res)
}