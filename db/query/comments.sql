-- CREATE TABLE comments (
--   "id" BIGSERIAL PRIMARY KEY,
--   "post_id" BIGINT NOT NULL,
--   "parent_comment_id" BIGINT NULL,
--   "commentor" VARCHAR(255) NOT NULL,
--   "comment_text" TEXT NOT NULL,
--   "created_at" TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,  
--   FOREIGN KEY ("post_id") REFERENCES posts("id"),
--   FOREIGN KEY ("commentor") REFERENCES users("username"),
--   FOREIGN KEY ("parent_comment_id") REFERENCES comments("id")
-- );

-- name: CreateComment :one
INSERT INTO comments (post_id, commentor, comment_text, parent_comment_id) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments WHERE id = $1 LIMIT 1;

-- name: GetComments :many
SELECT * FROM comments ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateComment :exec
UPDATE comments SET comment_text = $2 WHERE id = $1;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;