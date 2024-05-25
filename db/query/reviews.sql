-- name: CreateReview :one
INSERT INTO reviews (seller, reviewer, rating, content) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetReview :one
SELECT * FROM reviews WHERE id = $1 LIMIT 1;

-- name: GetReviews :many
SELECT * FROM reviews ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateReview :exec
UPDATE reviews SET seller = $2, reviewer = $3, rating = $4, content = $5 WHERE id = $1;

-- name: DeleteReview :exec
DELETE FROM reviews WHERE id = $1;
