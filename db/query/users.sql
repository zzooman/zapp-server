-- name: CreateUser :one
INSERT INTO users (username, password, email, phone) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users SET password = $2, phone = $3, email = $4, profile = $5 WHERE username = $1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;