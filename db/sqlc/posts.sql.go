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

const getPostWithAuthor = `-- name: GetPostWithAuthor :one
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username WHERE posts.id = $1 LIMIT 1
`

type GetPostWithAuthorRow struct {
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

func (q *Queries) GetPostWithAuthor(ctx context.Context, id int64) (GetPostWithAuthorRow, error) {
	row := q.db.QueryRow(ctx, getPostWithAuthor, id)
	var i GetPostWithAuthorRow
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
		&i.Email,
		&i.Phone,
		&i.Profile,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, author, title, content, medias, price, stock, views, created_at FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2
`

type GetPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPosts(ctx context.Context, arg GetPostsParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, getPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
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

const getPostsWithAuthorByQuery = `-- name: GetPostsWithAuthorByQuery :many
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username WHERE posts.title ILIKE '%' || $1 || '%' OR posts.content ILIKE '%' || $1 || '%' ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3
`

type GetPostsWithAuthorByQueryParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

type GetPostsWithAuthorByQueryRow struct {
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

func (q *Queries) GetPostsWithAuthorByQuery(ctx context.Context, arg GetPostsWithAuthorByQueryParams) ([]GetPostsWithAuthorByQueryRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithAuthorByQuery, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithAuthorByQueryRow{}
	for rows.Next() {
		var i GetPostsWithAuthorByQueryRow
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

const getPostsWithAuthorThatIBought = `-- name: GetPostsWithAuthorThatIBought :many
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username JOIN transactions ON posts.id = transactions.post_id WHERE transactions.buyer = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3
`

type GetPostsWithAuthorThatIBoughtParams struct {
	Buyer  string `json:"buyer"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetPostsWithAuthorThatIBoughtRow struct {
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

func (q *Queries) GetPostsWithAuthorThatIBought(ctx context.Context, arg GetPostsWithAuthorThatIBoughtParams) ([]GetPostsWithAuthorThatIBoughtRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithAuthorThatIBought, arg.Buyer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithAuthorThatIBoughtRow{}
	for rows.Next() {
		var i GetPostsWithAuthorThatIBoughtRow
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

const getPostsWithAuthorThatILiked = `-- name: GetPostsWithAuthorThatILiked :many
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username JOIN like_with_post ON posts.id = like_with_post.post_id WHERE like_with_post.username = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3
`

type GetPostsWithAuthorThatILikedParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type GetPostsWithAuthorThatILikedRow struct {
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

func (q *Queries) GetPostsWithAuthorThatILiked(ctx context.Context, arg GetPostsWithAuthorThatILikedParams) ([]GetPostsWithAuthorThatILikedRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithAuthorThatILiked, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithAuthorThatILikedRow{}
	for rows.Next() {
		var i GetPostsWithAuthorThatILikedRow
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

const getPostsWithAuthorThatISold = `-- name: GetPostsWithAuthorThatISold :many
SELECT posts.id, posts.author, posts.title, posts.content, posts.medias, posts.price, posts.stock, posts.views, posts.created_at, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username JOIN transactions ON posts.id = transactions.post_id WHERE transactions.seller = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3
`

type GetPostsWithAuthorThatISoldParams struct {
	Seller string `json:"seller"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

type GetPostsWithAuthorThatISoldRow struct {
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

func (q *Queries) GetPostsWithAuthorThatISold(ctx context.Context, arg GetPostsWithAuthorThatISoldParams) ([]GetPostsWithAuthorThatISoldRow, error) {
	rows, err := q.db.Query(ctx, getPostsWithAuthorThatISold, arg.Seller, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetPostsWithAuthorThatISoldRow{}
	for rows.Next() {
		var i GetPostsWithAuthorThatISoldRow
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
