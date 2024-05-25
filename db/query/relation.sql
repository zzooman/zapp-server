-- name: CreateLikeWithPost :one
INSERT INTO like_with_post (username, post_id) VALUES ($1, $2) RETURNING *;
-- name: GetLikeWithPost :one
SELECT * FROM like_with_post WHERE username = $1 AND post_id = $2 LIMIT 1;
-- name: DeleteLikeWithPost :exec
DELETE FROM like_with_post WHERE username = $1 AND post_id = $2;


-- name: CreateWishWithProduct :one
INSERT INTO wish_with_product (username, post_id) VALUES ($1, $2) RETURNING *;
-- name: GetWishWithProduct :one
SELECT * FROM wish_with_product WHERE username = $1 AND post_id = $2 LIMIT 1;
-- name: DeleteWishWithProduct :exec
DELETE FROM wish_with_product WHERE username = $1 AND post_id = $2;
