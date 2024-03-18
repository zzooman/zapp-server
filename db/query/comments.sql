-- name: CreateComment :one
INSERT INTO comments (post_id, commentor, content) VALUES ($1, $2, $3) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments WHERE id = $1 LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateComment :exec
UPDATE comments SET post_id = $2, commentor = $3, content = $4 WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;