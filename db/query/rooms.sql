-- name: CheckRoom :one
SELECT * FROM Rooms WHERE (user_a = $1 AND user_b = $2) OR (user_a = $2 AND user_b = $1) LIMIT 1;

-- name: CreateRoom :one
INSERT INTO Rooms (user_a, user_b) VALUES ($1, $2) RETURNING *;

-- name: GetRoom :one
SELECT * FROM Rooms WHERE id = $1 LIMIT 1;

-- name: GetRoomsByUser :many
SELECT * FROM Rooms WHERE user_a = $1 OR user_b = $1 ORDER BY id;


