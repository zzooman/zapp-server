// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: reviews.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createReview = `-- name: CreateReview :one
INSERT INTO reviews (product_id, reviewer, rating, content) VALUES ($1, $2, $3, $4) RETURNING id, product_id, reviewer, rating, medias, content, created_at
`

type CreateReviewParams struct {
	ProductID int64       `json:"product_id"`
	Reviewer  string      `json:"reviewer"`
	Rating    int32       `json:"rating"`
	Content   pgtype.Text `json:"content"`
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) (Review, error) {
	row := q.db.QueryRow(ctx, createReview,
		arg.ProductID,
		arg.Reviewer,
		arg.Rating,
		arg.Content,
	)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Reviewer,
		&i.Rating,
		&i.Medias,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM reviews WHERE id = $1
`

func (q *Queries) DeleteReview(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteReview, id)
	return err
}

const getReview = `-- name: GetReview :one
SELECT id, product_id, reviewer, rating, medias, content, created_at FROM reviews WHERE id = $1 LIMIT 1
`

func (q *Queries) GetReview(ctx context.Context, id int64) (Review, error) {
	row := q.db.QueryRow(ctx, getReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.Reviewer,
		&i.Rating,
		&i.Medias,
		&i.Content,
		&i.CreatedAt,
	)
	return i, err
}

const getReviews = `-- name: GetReviews :many
SELECT id, product_id, reviewer, rating, medias, content, created_at FROM reviews ORDER BY id LIMIT $1 OFFSET $2
`

type GetReviewsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetReviews(ctx context.Context, arg GetReviewsParams) ([]Review, error) {
	rows, err := q.db.Query(ctx, getReviews, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Review{}
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ID,
			&i.ProductID,
			&i.Reviewer,
			&i.Rating,
			&i.Medias,
			&i.Content,
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

const updateReview = `-- name: UpdateReview :exec
UPDATE reviews SET product_id = $2, reviewer = $3, rating = $4, content = $5 WHERE id = $1
`

type UpdateReviewParams struct {
	ID        int64       `json:"id"`
	ProductID int64       `json:"product_id"`
	Reviewer  string      `json:"reviewer"`
	Rating    int32       `json:"rating"`
	Content   pgtype.Text `json:"content"`
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.Exec(ctx, updateReview,
		arg.ID,
		arg.ProductID,
		arg.Reviewer,
		arg.Rating,
		arg.Content,
	)
	return err
}
