-- name: CreateOrder :one
INSERT INTO orders (product_id, buyer_id, quantity, price_at_order, status) VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders WHERE id = $1 LIMIT 1;

-- name: GetOrders :many
SELECT * FROM orders ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateOrder :exec
UPDATE orders SET product_id = $2, buyer_id = $3, quantity = $4, price_at_order = $5, status = $6 WHERE id = $1;

-- name: DeleteOrder :exec
DELETE FROM orders WHERE id = $1;