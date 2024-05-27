CREATE TABLE posts (
  "id" BIGSERIAL PRIMARY KEY,
  "author" VARCHAR(255) NOT NULL,  
  "title" VARCHAR(255) NOT NULL,
  "content" TEXT NOT NULL,
  "media" VARCHAR[],  
  "price" BIGINT NOT NULL,
  "stock" BIGINT NOT NULL,
  "views" BIGINT DEFAULT 0,
  "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,  
  FOREIGN KEY ("author") REFERENCES users("username")
);


-- name: CreatePost :one
INSERT INTO posts (author, title, content, price, stock, media, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: GetPost :one
SELECT * FROM posts WHERE id = $1 LIMIT 1;

-- name: GetPostsWithAuthor :many
SELECT posts.*, users.email, users.phone, users.profile FROM posts JOIN users ON posts.author = users.username ORDER BY posts.created_at DESC LIMIT $1 OFFSET $2;


-- name: UpdatePost :exec
UPDATE posts SET title = $2, content = $3, price = $4, stock = $5, media = $6 WHERE id = $1;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;


