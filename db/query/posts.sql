-- name: CreatePost :one
INSERT INTO posts (author, product_id, title, content, media, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdatePost :exec
UPDATE posts SET author = $2, product_id = $3, title = $4, content = $5, media = $6 WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;


