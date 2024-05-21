-- name: CreatePayment :one
INSERT INTO payments (transaction_id, payment_status, payment_method, payment_amount) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments WHERE payment_id = $1 LIMIT 1;

-- name: GetPayments :many
SELECT * FROM payments WHERE transaction_id = $1 ORDER BY payment_id LIMIT $2 OFFSET $3;

-- name: UpdatePaymentStatus :one
UPDATE payments SET payment_status = $2 WHERE payment_id = $1 RETURNING *;

-- name: DeletePayment :exec
DELETE FROM payments WHERE payment_id = $1;