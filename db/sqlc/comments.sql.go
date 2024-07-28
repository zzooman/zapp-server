// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: comments.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (post_id, commentor, comment_text, parent_comment_id) VALUES ($1, $2, $3, $4) RETURNING id, post_id, parent_comment_id, commentor, comment_text, created_at
`

type CreateCommentParams struct {
	PostID          int64       `json:"post_id"`
	Commentor       string      `json:"commentor"`
	CommentText     string      `json:"comment_text"`
	ParentCommentID pgtype.Int8 `json:"parent_comment_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.PostID,
		arg.Commentor,
		arg.CommentText,
		arg.ParentCommentID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.ParentCommentID,
		&i.Commentor,
		&i.CommentText,
		&i.CreatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1
`

func (q *Queries) DeleteComment(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteComment, id)
	return err
}

const getComment = `-- name: GetComment :one
SELECT id, post_id, parent_comment_id, commentor, comment_text, created_at FROM comments WHERE id = $1 LIMIT 1
`

func (q *Queries) GetComment(ctx context.Context, id int64) (Comment, error) {
	row := q.db.QueryRow(ctx, getComment, id)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.PostID,
		&i.ParentCommentID,
		&i.Commentor,
		&i.CommentText,
		&i.CreatedAt,
	)
	return i, err
}

const getComments = `-- name: GetComments :many
SELECT id, post_id, parent_comment_id, commentor, comment_text, created_at FROM comments ORDER BY id LIMIT $1 OFFSET $2
`

type GetCommentsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetComments(ctx context.Context, arg GetCommentsParams) ([]Comment, error) {
	rows, err := q.db.Query(ctx, getComments, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.PostID,
			&i.ParentCommentID,
			&i.Commentor,
			&i.CommentText,
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

const updateComment = `-- name: UpdateComment :exec
UPDATE comments SET comment_text = $2 WHERE id = $1
`

type UpdateCommentParams struct {
	ID          int64  `json:"id"`
	CommentText string `json:"comment_text"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) error {
	_, err := q.db.Exec(ctx, updateComment, arg.ID, arg.CommentText)
	return err
}
