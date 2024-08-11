-- name: CheckRoom :one
SELECT * FROM Rooms WHERE (host = $1 AND guest = $2) OR (host = $2 AND guest = $1) LIMIT 1;

-- name: CreateRoom :one
INSERT INTO Rooms (host, guest, product_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetRoom :one
SELECT * FROM Rooms WHERE id = $1 LIMIT 1;

-- name: GetRoomsByUser :many
SELECT * FROM Rooms WHERE host = $1 OR guest = $1 ORDER BY id;

-- name: GetChatsRoomByUser :many
SELECT * FROM Rooms WHERE (host = $1 OR guest = $1) AND product_id IS NULL;

-- name: GetSellRoomByUser :many
SELECT * FROM Rooms WHERE host = $1 AND product_id IS NOT NULL;

-- name: GetBuyRoomByUser :many
SELECT * FROM Rooms WHERE guest = $1 AND product_id IS NOT NUll;

-- name: DeleteRoom :one
DELETE FROM Rooms WHERE id = $1 RETURNING *;



