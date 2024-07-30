// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CheckRoom(ctx context.Context, arg CheckRoomParams) (Room, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error)
	CreateLikeWithFeed(ctx context.Context, arg CreateLikeWithFeedParams) (LikeWithFeed, error)
	CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error)
	CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error)
	CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error)
	CreateRoom(ctx context.Context, arg CreateRoomParams) (Room, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateWishWithProduct(ctx context.Context, arg CreateWishWithProductParams) (WishWithProduct, error)
	DeleteComment(ctx context.Context, id int64) error
	DeleteFeed(ctx context.Context, id int64) error
	DeleteLikeWithFeed(ctx context.Context, arg DeleteLikeWithFeedParams) error
	DeleteMessage(ctx context.Context, id int64) (Message, error)
	DeleteProduct(ctx context.Context, id int64) error
	DeleteReview(ctx context.Context, id int64) error
	DeleteRoom(ctx context.Context, id int64) (Room, error)
	DeleteTransaction(ctx context.Context, transactionID int64) error
	DeleteUser(ctx context.Context, username string) error
	DeleteWishWithProduct(ctx context.Context, arg DeleteWishWithProductParams) error
	GetBuyerTransactions(ctx context.Context, arg GetBuyerTransactionsParams) ([]Transaction, error)
	GetComment(ctx context.Context, id int64) (Comment, error)
	GetComments(ctx context.Context, arg GetCommentsParams) ([]Comment, error)
	GetFeed(ctx context.Context, id int64) (Feed, error)
	GetFeeds(ctx context.Context, arg GetFeedsParams) ([]Feed, error)
	GetFeedsWithAuthor(ctx context.Context, arg GetFeedsWithAuthorParams) ([]GetFeedsWithAuthorRow, error)
	GetFeedsWithAuthorByQuery(ctx context.Context, arg GetFeedsWithAuthorByQueryParams) ([]GetFeedsWithAuthorByQueryRow, error)
	GetFeedsWithAuthorThatIBought(ctx context.Context, arg GetFeedsWithAuthorThatIBoughtParams) ([]GetFeedsWithAuthorThatIBoughtRow, error)
	GetFeedsWithAuthorThatISold(ctx context.Context, arg GetFeedsWithAuthorThatISoldParams) ([]GetFeedsWithAuthorThatISoldRow, error)
	GetFeedsWithAuthorThatIWished(ctx context.Context, arg GetFeedsWithAuthorThatIWishedParams) ([]GetFeedsWithAuthorThatIWishedRow, error)
	GetLastMessage(ctx context.Context, roomID int64) (Message, error)
	GetLikeWithFeed(ctx context.Context, arg GetLikeWithFeedParams) (LikeWithFeed, error)
	GetMessagesByRoom(ctx context.Context, roomID int64) ([]Message, error)
	GetProduct(ctx context.Context, id int64) (Product, error)
	GetProductWithAuthor(ctx context.Context, id int64) (GetProductWithAuthorRow, error)
	GetProductWithSellor(ctx context.Context, id int64) (GetProductWithSellorRow, error)
	GetProducts(ctx context.Context, arg GetProductsParams) ([]Product, error)
	GetProductsWithSeller(ctx context.Context, arg GetProductsWithSellerParams) ([]GetProductsWithSellerRow, error)
	GetProductsWithSellerByQuery(ctx context.Context, arg GetProductsWithSellerByQueryParams) ([]GetProductsWithSellerByQueryRow, error)
	GetProductsWithSellerThatIBought(ctx context.Context, arg GetProductsWithSellerThatIBoughtParams) ([]GetProductsWithSellerThatIBoughtRow, error)
	GetProductsWithSellerThatILiked(ctx context.Context, arg GetProductsWithSellerThatILikedParams) ([]GetProductsWithSellerThatILikedRow, error)
	GetProductsWithSellerThatISold(ctx context.Context, arg GetProductsWithSellerThatISoldParams) ([]GetProductsWithSellerThatISoldRow, error)
	GetReview(ctx context.Context, id int64) (Review, error)
	GetReviews(ctx context.Context, arg GetReviewsParams) ([]Review, error)
	GetRoom(ctx context.Context, id int64) (Room, error)
	GetRoomsByUser(ctx context.Context, userA string) ([]Room, error)
	GetSearchCount(ctx context.Context, searchText string) (SearchCount, error)
	GetSellerTransactions(ctx context.Context, arg GetSellerTransactionsParams) ([]Transaction, error)
	GetTransaction(ctx context.Context, transactionID int64) (Transaction, error)
	GetUser(ctx context.Context, username string) (User, error)
	GetWishWithProduct(ctx context.Context, arg GetWishWithProductParams) (WishWithProduct, error)
	HotSearchTexts(ctx context.Context) ([]SearchCount, error)
	ReadMessage(ctx context.Context, id int64) error
	UnreadMessageCount(ctx context.Context, sender string) (int64, error)
	UpdateComment(ctx context.Context, arg UpdateCommentParams) error
	UpdateFeed(ctx context.Context, arg UpdateFeedParams) error
	UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error)
	UpdateProduct(ctx context.Context, arg UpdateProductParams) error
	UpdateReview(ctx context.Context, arg UpdateReviewParams) error
	UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transaction, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	UpsertSearchCount(ctx context.Context, searchText string) (SearchCount, error)
}

var _ Querier = (*Queries)(nil)
