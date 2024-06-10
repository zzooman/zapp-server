-- name: CreateMessage :one
INSERT INTO Messages (room_id, sender, message) VALUES ($1, $2, $3) RETURNING *;

-- name: GetMessagesByRoom :many
SELECT * FROM Messages WHERE room_id = $1 ORDER BY id;

-- name: UpdateMessage :one
UPDATE Messages SET message = $2 WHERE id = $1 RETURNING *;

-- name: DeleteMessage :one
DELETE FROM Messages WHERE id = $1 RETURNING *;

-- name: GetLastMessage :one
SELECT * FROM Messages WHERE room_id = $1 ORDER BY id DESC LIMIT 1;