-- name: GetSearchCount :one
SELECT * FROM search_count WHERE search_text = $1 LIMIT 1;

-- name: UpsertSearchCount :one
INSERT INTO search_count (search_text, count) 
VALUES ($1, 1) 
ON CONFLICT (search_text) 
DO UPDATE SET count = search_count.count + 1 
RETURNING *;

