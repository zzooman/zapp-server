// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: relation.sql

package db

import (
	"context"
)

const createLikeWithFeed = `-- name: CreateLikeWithFeed :one
INSERT INTO like_with_feed (username, feed_id) VALUES ($1, $2) RETURNING username, feed_id
`

type CreateLikeWithFeedParams struct {
	Username string `json:"username"`
	FeedID   int64  `json:"feed_id"`
}

func (q *Queries) CreateLikeWithFeed(ctx context.Context, arg CreateLikeWithFeedParams) (LikeWithFeed, error) {
	row := q.db.QueryRow(ctx, createLikeWithFeed, arg.Username, arg.FeedID)
	var i LikeWithFeed
	err := row.Scan(&i.Username, &i.FeedID)
	return i, err
}

const createWishWithProduct = `-- name: CreateWishWithProduct :one
INSERT INTO wish_with_product (username, product_id) VALUES ($1, $2) RETURNING username, product_id
`

type CreateWishWithProductParams struct {
	Username  string `json:"username"`
	ProductID int64  `json:"product_id"`
}

func (q *Queries) CreateWishWithProduct(ctx context.Context, arg CreateWishWithProductParams) (WishWithProduct, error) {
	row := q.db.QueryRow(ctx, createWishWithProduct, arg.Username, arg.ProductID)
	var i WishWithProduct
	err := row.Scan(&i.Username, &i.ProductID)
	return i, err
}

const deleteLikeWithFeed = `-- name: DeleteLikeWithFeed :exec
DELETE FROM like_with_feed WHERE username = $1 AND feed_id = $2
`

type DeleteLikeWithFeedParams struct {
	Username string `json:"username"`
	FeedID   int64  `json:"feed_id"`
}

func (q *Queries) DeleteLikeWithFeed(ctx context.Context, arg DeleteLikeWithFeedParams) error {
	_, err := q.db.Exec(ctx, deleteLikeWithFeed, arg.Username, arg.FeedID)
	return err
}

const deleteWishWithProduct = `-- name: DeleteWishWithProduct :exec
DELETE FROM wish_with_product WHERE username = $1 AND product_id = $2
`

type DeleteWishWithProductParams struct {
	Username  string `json:"username"`
	ProductID int64  `json:"product_id"`
}

func (q *Queries) DeleteWishWithProduct(ctx context.Context, arg DeleteWishWithProductParams) error {
	_, err := q.db.Exec(ctx, deleteWishWithProduct, arg.Username, arg.ProductID)
	return err
}

const getLikeWithFeed = `-- name: GetLikeWithFeed :one
SELECT username, feed_id FROM like_with_feed WHERE username = $1 AND feed_id = $2 LIMIT 1
`

type GetLikeWithFeedParams struct {
	Username string `json:"username"`
	FeedID   int64  `json:"feed_id"`
}

func (q *Queries) GetLikeWithFeed(ctx context.Context, arg GetLikeWithFeedParams) (LikeWithFeed, error) {
	row := q.db.QueryRow(ctx, getLikeWithFeed, arg.Username, arg.FeedID)
	var i LikeWithFeed
	err := row.Scan(&i.Username, &i.FeedID)
	return i, err
}

const getWishWithProduct = `-- name: GetWishWithProduct :one
SELECT username, product_id FROM wish_with_product WHERE username = $1 AND product_id = $2 LIMIT 1
`

type GetWishWithProductParams struct {
	Username  string `json:"username"`
	ProductID int64  `json:"product_id"`
}

func (q *Queries) GetWishWithProduct(ctx context.Context, arg GetWishWithProductParams) (WishWithProduct, error) {
	row := q.db.QueryRow(ctx, getWishWithProduct, arg.Username, arg.ProductID)
	var i WishWithProduct
	err := row.Scan(&i.Username, &i.ProductID)
	return i, err
}
