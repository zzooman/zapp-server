-- name: CreateProduct :one
INSERT INTO products (seller, name, description, price, stock, images) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateProduct :exec
UPDATE products SET seller = $2, name = $3, description = $4, price = $5, stock = $6, images = $7 WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;