-- name: CreatePost :one
INSERT INTO posts (author, content, medias, created_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPosts :many
SELECT * FROM posts ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetProductWithAuthor :one
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.seller = users.username WHERE posts.id = $1 LIMIT 1;

-- name: GetPostsWithAuthor :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.seller = users.username ORDER BY posts.created_at DESC LIMIT $1 OFFSET $2;

-- name: GetPostsWithAuthorByQuery :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.seller = users.username WHERE posts.content ILIKE '%' || $1 || '%' ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetPostsWithAuthorThatIWished :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.seller = users.username JOIN wish_with_product ON posts.id = wish_with_product.product_id WHERE wish_with_product.username = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetPostsWithAuthorThatIBought :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.seller = users.username JOIN transactions ON posts.id = transactions.product_id WHERE transactions.buyer = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetPostsWithAuthorThatISold :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username JOIN transactions ON posts.id = transactions.product_id WHERE transactions.seller = $1 ORDER BY posts.created_at DESC LIMIT $2 OFFSET $3;

-- name: UpdatePost :exec
UPDATE posts SET content = $2, medias = $3 WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;









