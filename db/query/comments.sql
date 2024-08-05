-- name: CreateComment :one
INSERT INTO comments (feed_id, commentor, comment_text, parent_comment_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments WHERE id = $1 LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments WHERE feed_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateComment :exec
UPDATE comments SET comment_text = $2 WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;

-- name: GetCountOfComments :one
SELECT COUNT(*) FROM comments WHERE feed_id = $1;