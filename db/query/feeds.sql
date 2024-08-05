-- name: CreateFeed :one
INSERT INTO feeds (author, content, medias, created_at) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetFeed :one
SELECT * FROM feeds WHERE id = $1 LIMIT 1;

-- name: GetFeedWithAuthor :one
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username WHERE feeds.id = $1 LIMIT 1;

-- name: GetFeeds :many
SELECT * FROM feeds ORDER BY created_at DESC LIMIT $1 OFFSET $2;

-- name: GetFeedsWithAuthor :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username ORDER BY feeds.created_at DESC LIMIT $1 OFFSET $2;

-- name: GetFeedsWithAuthorByQuery :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username WHERE feeds.content ILIKE '%' || $1 || '%' ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: GetFeedsWithAuthorThatILiked :many
SELECT feeds.*, users.email, users.phone, users.profile FROM feeds JOIN users ON feeds.author = users.username JOIN like_with_feed ON feeds.id = like_with_feed.feed_id WHERE like_with_feed.username = $1 ORDER BY feeds.created_at DESC LIMIT $2 OFFSET $3;

-- name: UpdateFeed :exec
UPDATE feeds SET content = $2, medias = $3 WHERE id = $1;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE id = $1;









