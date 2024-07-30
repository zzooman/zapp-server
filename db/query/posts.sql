-- name: CreateFeed :one
INSERT INTO feeds (author, content, medias, created_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1 LIMIT 1;

-- name: GetFeeds :many
SELECT * FROM feeds ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetProductWithAuthor :one
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.seller = users.username WHERE feeds.id = $1 LIMIT 1;

-- name: GetFeedsWithAuthor :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.seller = users.username ORDER BY feeds.created_at DESC LIMIT $1 OFFSET $2;

-- name: GetFeedsWithAuthorByQuery :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.seller = users.username WHERE feeds.content ILIKE '%' || $1 || '%' ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetFeedsWithAuthorThatIWished :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.seller = users.username JOIN wish_with_product ON feeds.id = wish_with_product.product_id WHERE wish_with_product.username = $1 ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetFeedsWithAuthorThatIBought :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.seller = users.username JOIN transactions ON feeds.id = transactions.product_id WHERE transactions.buyer = $1 ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetFeedsWithAuthorThatISold :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username JOIN transactions ON feeds.id = transactions.product_id WHERE transactions.seller = $1 ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: UpdateFeed :exec
UPDATE feeds SET content = $2, medias = $3 WHERE id = $1;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;









