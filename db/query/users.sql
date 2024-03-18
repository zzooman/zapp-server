-- name: CreateUser :one
INSERT INTO users (username, password, phone, email, location) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: UpdateUser :exec
UPDATE users SET username = $1, password = $2, phone = $3, email = $4, location = $5 WHERE username = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1;