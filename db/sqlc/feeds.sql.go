// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feeds.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (author, content, medias, created_at) VALUES ($1, $2, $3, $4) RETURNING id, author, content, medias, views, created_at
`

type CreateFeedParams struct {
	Author    string             `json:"author"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRow(ctx, createFeed,
		arg.Author,
		arg.Content,
		arg.Medias,
		arg.CreatedAt,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Content,
		&i.Medias,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const deleteFeed = `-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1
`

func (q *Queries) DeleteFeed(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteFeed, id)
	return err
}

const getFeed = `-- name: GetFeed :one
SELECT id, author, content, medias, views, created_at FROM feeds WHERE id = $1 LIMIT 1
`

func (q *Queries) GetFeed(ctx context.Context, id int64) (Feed, error) {
	row := q.db.QueryRow(ctx, getFeed, id)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Content,
		&i.Medias,
		&i.Views,
		&i.CreatedAt,
	)
	return i, err
}

const getFeedWithAuthor = `-- name: GetFeedWithAuthor :one
SELECT feeds.id, feeds.author, feeds.content, feeds.medias, feeds.views, feeds.created_at, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username WHERE feeds.id = $1 LIMIT 1
`

type GetFeedWithAuthorRow struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetFeedWithAuthor(ctx context.Context, id int64) (GetFeedWithAuthorRow, error) {
	row := q.db.QueryRow(ctx, getFeedWithAuthor, id)
	var i GetFeedWithAuthorRow
	err := row.Scan(
		&i.ID,
		&i.Author,
		&i.Content,
		&i.Medias,
		&i.Views,
		&i.CreatedAt,
		&i.Email,
		&i.Phone,
		&i.Profile,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, author, content, medias, views, created_at FROM feeds ORDER BY created_at DESC LIMIT $1 OFFSET $2
`

type GetFeedsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetFeeds(ctx context.Context, arg GetFeedsParams) ([]Feed, error) {
	rows, err := q.db.Query(ctx, getFeeds, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Feed{}
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Content,
			&i.Medias,
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

const getFeedsWithAuthor = `-- name: GetFeedsWithAuthor :many
SELECT feeds.id, feeds.author, feeds.content, feeds.medias, feeds.views, feeds.created_at, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username ORDER BY feeds.created_at DESC LIMIT $1 OFFSET $2
`

type GetFeedsWithAuthorParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetFeedsWithAuthorRow struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetFeedsWithAuthor(ctx context.Context, arg GetFeedsWithAuthorParams) ([]GetFeedsWithAuthorRow, error) {
	rows, err := q.db.Query(ctx, getFeedsWithAuthor, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFeedsWithAuthorRow{}
	for rows.Next() {
		var i GetFeedsWithAuthorRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Content,
			&i.Medias,
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

const getFeedsWithAuthorByQuery = `-- name: GetFeedsWithAuthorByQuery :many
SELECT feeds.id, feeds.author, feeds.content, feeds.medias, feeds.views, feeds.created_at, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username WHERE feeds.content ILIKE '%' || $1 || '%' ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3
`

type GetFeedsWithAuthorByQueryParams struct {
	Column1 pgtype.Text `json:"column_1"`
	Limit   int32       `json:"limit"`
	Offset  int32       `json:"offset"`
}

type GetFeedsWithAuthorByQueryRow struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetFeedsWithAuthorByQuery(ctx context.Context, arg GetFeedsWithAuthorByQueryParams) ([]GetFeedsWithAuthorByQueryRow, error) {
	rows, err := q.db.Query(ctx, getFeedsWithAuthorByQuery, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFeedsWithAuthorByQueryRow{}
	for rows.Next() {
		var i GetFeedsWithAuthorByQueryRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Content,
			&i.Medias,
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

const getFeedsWithAuthorThatILiked = `-- name: GetFeedsWithAuthorThatILiked :many
SELECT feeds.id, feeds.author, feeds.content, feeds.medias, feeds.views, feeds.created_at, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username JOIN like_with_feed ON feeds.id = like_with_feed.feed_id WHERE like_with_feed.username = $1 ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3
`

type GetFeedsWithAuthorThatILikedParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

type GetFeedsWithAuthorThatILikedRow struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Content   string             `json:"content"`
	Medias    []string           `json:"medias"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Email     string             `json:"email"`
	Phone     pgtype.Text        `json:"phone"`
	Profile   pgtype.Text        `json:"profile"`
}

func (q *Queries) GetFeedsWithAuthorThatILiked(ctx context.Context, arg GetFeedsWithAuthorThatILikedParams) ([]GetFeedsWithAuthorThatILikedRow, error) {
	rows, err := q.db.Query(ctx, getFeedsWithAuthorThatILiked, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetFeedsWithAuthorThatILikedRow{}
	for rows.Next() {
		var i GetFeedsWithAuthorThatILikedRow
		if err := rows.Scan(
			&i.ID,
			&i.Author,
			&i.Content,
			&i.Medias,
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

const updateFeed = `-- name: UpdateFeed :exec
UPDATE feeds SET content = $2, medias = $3 WHERE id = $1
`

type UpdateFeedParams struct {
	ID      int64    `json:"id"`
	Content string   `json:"content"`
	Medias  []string `json:"medias"`
}

func (q *Queries) UpdateFeed(ctx context.Context, arg UpdateFeedParams) error {
	_, err := q.db.Exec(ctx, updateFeed, arg.ID, arg.Content, arg.Medias)
	return err
}

const viewFeed = `-- name: ViewFeed :exec
UPDATE feeds SET views = views + 1 WHERE id = $1
`

func (q *Queries) ViewFeed(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, viewFeed, id)
	return err
}
