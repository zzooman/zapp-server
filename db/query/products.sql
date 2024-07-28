-- name: CreateProduct :one
INSERT INTO products (seller, title, content, price, medias, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products WHERE id = $1 LIMIT 1;

-- name: GetProducts :many
SELECT * FROM products ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetProductWithSeller :one
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username WHERE products.id = $1 LIMIT 1;

-- name: GetProductsWithSeller :many
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username ORDER BY products.created_at DESC LIMIT $1 OFFSET $2;

-- name: GetProductsWithSellerByQuery :many
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username WHERE products.title ILIKE '%' || $1 || '%' OR products.content ILIKE '%' || $1 || '%' ORDER BY products.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetProductsWithSellerThatIWished :many
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username JOIN wish_with_product ON products.id = wish_with_product.product_id WHERE wish_with_product.username = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetProductsWithSellerThatIBought :many
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON products.seller = users.username JOIN transactions ON products.id = transactions.product_id WHERE transactions.buyer = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetProductsWithSellerThatISold :many
SELECT products.*, users.email, users.phone, users.profile FROM products JOIN users ON posts.author = users.username JOIN transactions ON products.id = transactions.product_id WHERE transactions.seller = $1 ORDER BY products.created_at DESC LIMIT $2 OFFSET $3;

-- name: UpdateProduct :exec
UPDATE products SET title = $2, content = $3, price = $4, medias = $5 WHERE id = $1;

-- name: DeleteProduct :exec
DELETE FROM products WHERE id = $1;


