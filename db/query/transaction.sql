-- name: CreateTransaction :one
INSERT INTO transactions (product_id, buyer, seller) VALUES ($1, $2, $3) RETURNING *;

-- name: GetTransaction :one
SELECT * FROM transactions WHERE transaction_id = $1 LIMIT 1;

-- name: GetBuyerTransactions :many
SELECT * FROM transactions WHERE buyer = $1 ORDER BY transaction_id LIMIT $2 OFFSET $3;

-- name: GetSellerTransactions :many
SELECT * FROM transactions WHERE seller = $1 ORDER BY transaction_id LIMIT $2 OFFSET $3;

-- name: UpdateTransactionStatus :one
UPDATE transactions SET status = $2 WHERE transaction_id = $1 RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE transaction_id = $1;

