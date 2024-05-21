// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: payments.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (transaction_id, payment_status, payment_method, payment_amount) VALUES ($1, $2, $3, $4) RETURNING payment_id, transaction_id, payment_status, payment_method, payment_date, payment_amount
`

type CreatePaymentParams struct {
	TransactionID int64          `json:"transaction_id"`
	PaymentStatus pgtype.Text    `json:"payment_status"`
	PaymentMethod string         `json:"payment_method"`
	PaymentAmount pgtype.Numeric `json:"payment_amount"`
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRow(ctx, createPayment,
		arg.TransactionID,
		arg.PaymentStatus,
		arg.PaymentMethod,
		arg.PaymentAmount,
	)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.TransactionID,
		&i.PaymentStatus,
		&i.PaymentMethod,
		&i.PaymentDate,
		&i.PaymentAmount,
	)
	return i, err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments WHERE payment_id = $1
`

func (q *Queries) DeletePayment(ctx context.Context, paymentID int64) error {
	_, err := q.db.Exec(ctx, deletePayment, paymentID)
	return err
}

const getPayment = `-- name: GetPayment :one
SELECT payment_id, transaction_id, payment_status, payment_method, payment_date, payment_amount FROM payments WHERE payment_id = $1 LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, paymentID int64) (Payment, error) {
	row := q.db.QueryRow(ctx, getPayment, paymentID)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.TransactionID,
		&i.PaymentStatus,
		&i.PaymentMethod,
		&i.PaymentDate,
		&i.PaymentAmount,
	)
	return i, err
}

const getPayments = `-- name: GetPayments :many
SELECT payment_id, transaction_id, payment_status, payment_method, payment_date, payment_amount FROM payments WHERE transaction_id = $1 ORDER BY payment_id LIMIT $2 OFFSET $3
`

type GetPaymentsParams struct {
	TransactionID int64 `json:"transaction_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) GetPayments(ctx context.Context, arg GetPaymentsParams) ([]Payment, error) {
	rows, err := q.db.Query(ctx, getPayments, arg.TransactionID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Payment{}
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.PaymentID,
			&i.TransactionID,
			&i.PaymentStatus,
			&i.PaymentMethod,
			&i.PaymentDate,
			&i.PaymentAmount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePaymentStatus = `-- name: UpdatePaymentStatus :one
UPDATE payments SET payment_status = $2 WHERE payment_id = $1 RETURNING payment_id, transaction_id, payment_status, payment_method, payment_date, payment_amount
`

type UpdatePaymentStatusParams struct {
	PaymentID     int64       `json:"payment_id"`
	PaymentStatus pgtype.Text `json:"payment_status"`
}

func (q *Queries) UpdatePaymentStatus(ctx context.Context, arg UpdatePaymentStatusParams) (Payment, error) {
	row := q.db.QueryRow(ctx, updatePaymentStatus, arg.PaymentID, arg.PaymentStatus)
	var i Payment
	err := row.Scan(
		&i.PaymentID,
		&i.TransactionID,
		&i.PaymentStatus,
		&i.PaymentMethod,
		&i.PaymentDate,
		&i.PaymentAmount,
	)
	return i, err
}
