-- name: CreateUser :one
INSERT INTO users (username, password, phone, email, location) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE id = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users SET username = $2, password = $3, phone = $4, email = $5, location = $6 WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;