// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: posts.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (author, title, content, price, stock, medias, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, author, title, content, medias, price, stock, views, created_at
`

type CreatePostParams struct {
	Author    string             `json:"author"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Price     int64              `json:"price"`
	Stock     int64              `json:"stock"`
	Medias    []string           `json:"medias"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.Author,
		arg.Title,
		arg.Content,
		arg.Price,
		arg.Stock,
		arg.Medias,
		arg.CreatedAt,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.Content,
		&i.Medias,
		&i.Price,
		&i.Stock,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, author, title, content, medias, price, stock, views, created_at FROM posts WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int64) (Post, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Title,
		&i.Content,
		&i.Medias,
		&i.Price,
		&i.Stock,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const getPostsWithAuthor = `-- name: GetPostsWithAuthor :many
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username ORDER BY posts.created_at DESC LIMIT $1 OFFSET $2
`

type GetPostsWithAuthorParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetPostsWithAuthorRow struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Price     int64              `json:"price"`
	Stock     int64              `json:"stock"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetPostsWithAuthor(ctx context.Context, arg GetPostsWithAuthorParams) ([]GetPostsWithAuthorRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithAuthor, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithAuthorRow{}
	for rows.Next() {
		var i GetPostsWithAuthorRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Title,
			&i.Content,
			&i.Medias,
			&i.Price,
			&i.Stock,
			&i.Views,
			&i.CreatedAt,
			&i.Email,
			&i.Phone,
			&i.Profile,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, price = $4, stock = $5, medias = $6 WHERE id = $1
`

type UpdatePostParams struct {
	ID      int64    `json:"id"`
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Price   int64    `json:"price"`
	Stock   int64    `json:"stock"`
	Medias  []string `json:"medias"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) error {
	_, err := q.db.Exec(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.Price,
		arg.Stock,
		arg.Medias,
	)
	return err
}
