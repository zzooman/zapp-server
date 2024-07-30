-- name: CreateLikeWithFeed :one
INSERT INTO like_with_feed (username, feed_id) VALUES ($1, $2) RETURNING *;
-- name: GetLikeWithFeed :one
SELECT * FROM like_with_feed WHERE username = $1 AND feed_id = $2 LIMIT 1;
-- name: DeleteLikeWithFeed :exec
DELETE FROM like_with_feed WHERE username = $1 AND feed_id = $2;


-- name: CreateWishWithProduct :one
INSERT INTO wish_with_product (username, product_id) VALUES ($1, $2) RETURNING *;
-- name: GetWishWithProduct :one
SELECT * FROM wish_with_product WHERE username = $1 AND product_id = $2 LIMIT 1;
-- name: DeleteWishWithProduct :exec
DELETE FROM wish_with_product WHERE username = $1 AND product_id = $2;
