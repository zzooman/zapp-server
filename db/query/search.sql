-- name: GetSearchCount :exec
SELECT count FROM search_count WHERE search_text = $1;

-- name: CreateSearchCount :exec
INSERT INTO search_count (search_text) VALUES ($1);

-- name: IncreceSearchCount :exec
UPDATE search_count SET count = count + 1 WHERE search_text = $1;



