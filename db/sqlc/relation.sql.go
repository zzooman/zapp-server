// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: relation.sql

package db

import (
	"context"
)

const createLikeWithPost = `-- name: CreateLikeWithPost :one
INSERT INTO like_with_post (username, post_id) VALUES ($1, $2) RETURNING username, post_id
`

type CreateLikeWithPostParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) CreateLikeWithPost(ctx context.Context, arg CreateLikeWithPostParams) (LikeWithPost, error) {
	row := q.db.QueryRow(ctx, createLikeWithPost, arg.Username, arg.PostID)
	var i LikeWithPost
	err := row.Scan(&i.Username, &i.PostID)
	return i, err
}

const createWishWithProduct = `-- name: CreateWishWithProduct :one
INSERT INTO wish_with_product (username, post_id) VALUES ($1, $2) RETURNING username, post_id
`

type CreateWishWithProductParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) CreateWishWithProduct(ctx context.Context, arg CreateWishWithProductParams) (WishWithProduct, error) {
	row := q.db.QueryRow(ctx, createWishWithProduct, arg.Username, arg.PostID)
	var i WishWithProduct
	err := row.Scan(&i.Username, &i.PostID)
	return i, err
}

const deleteLikeWithPost = `-- name: DeleteLikeWithPost :exec
DELETE FROM like_with_post WHERE username = $1 AND post_id = $2
`

type DeleteLikeWithPostParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) DeleteLikeWithPost(ctx context.Context, arg DeleteLikeWithPostParams) error {
	_, err := q.db.Exec(ctx, deleteLikeWithPost, arg.Username, arg.PostID)
	return err
}

const deleteWishWithProduct = `-- name: DeleteWishWithProduct :exec
DELETE FROM wish_with_product WHERE username = $1 AND post_id = $2
`

type DeleteWishWithProductParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) DeleteWishWithProduct(ctx context.Context, arg DeleteWishWithProductParams) error {
	_, err := q.db.Exec(ctx, deleteWishWithProduct, arg.Username, arg.PostID)
	return err
}

const getLikeWithPost = `-- name: GetLikeWithPost :one
SELECT username, post_id FROM like_with_post WHERE username = $1 AND post_id = $2 LIMIT 1
`

type GetLikeWithPostParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) GetLikeWithPost(ctx context.Context, arg GetLikeWithPostParams) (LikeWithPost, error) {
	row := q.db.QueryRow(ctx, getLikeWithPost, arg.Username, arg.PostID)
	var i LikeWithPost
	err := row.Scan(&i.Username, &i.PostID)
	return i, err
}

const getWishWithProduct = `-- name: GetWishWithProduct :one
SELECT username, post_id FROM wish_with_product WHERE username = $1 AND post_id = $2 LIMIT 1
`

type GetWishWithProductParams struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

func (q *Queries) GetWishWithProduct(ctx context.Context, arg GetWishWithProductParams) (WishWithProduct, error) {
	row := q.db.QueryRow(ctx, getWishWithProduct, arg.Username, arg.PostID)
	var i WishWithProduct
	err := row.Scan(&i.Username, &i.PostID)
	return i, err
}
