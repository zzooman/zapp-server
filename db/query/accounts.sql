-- name: CreateAccount :one
INSERT INTO accounts (owner, bank_name, account_number, account_holder_name) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
-- SELECT * FROM accounts WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: GetAccounts :many
SELECT * FROM accounts WHERE owner = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: UpdateAccount :one
UPDATE accounts SET bank_name = $2, account_number = $3, account_holder_name = $4 WHERE id = $1 RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE id = $1;
