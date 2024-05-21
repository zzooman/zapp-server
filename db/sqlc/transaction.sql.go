// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: transaction.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTransaction = `-- name: CreateTransaction :one
INSERT INTO transactions (product_id, buyer, seller, total_amount) VALUES ($1, $2, $3, $4) RETURNING transaction_id, product_id, buyer, seller, status, total_amount, created_at
`

type CreateTransactionParams struct {
	ProductID   int64          `json:"product_id"`
	Buyer       string         `json:"buyer"`
	Seller      string         `json:"seller"`
	TotalAmount pgtype.Numeric `json:"total_amount"`
}

func (q *Queries) CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, createTransaction,
		arg.ProductID,
		arg.Buyer,
		arg.Seller,
		arg.TotalAmount,
	)
	var i Transaction
	err := row.Scan(
		&i.TransactionID,
		&i.ProductID,
		&i.Buyer,
		&i.Seller,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransaction = `-- name: DeleteTransaction :exec
DELETE FROM transactions WHERE transaction_id = $1
`

func (q *Queries) DeleteTransaction(ctx context.Context, transactionID int64) error {
	_, err := q.db.Exec(ctx, deleteTransaction, transactionID)
	return err
}

const getBuyerTransactions = `-- name: GetBuyerTransactions :many
SELECT transaction_id, product_id, buyer, seller, status, total_amount, created_at FROM transactions WHERE buyer = $1 ORDER BY transaction_id LIMIT $2 OFFSET $3
`

type GetBuyerTransactionsParams struct {
	Buyer  string `json:"buyer"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) GetBuyerTransactions(ctx context.Context, arg GetBuyerTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getBuyerTransactions, arg.Buyer, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionID,
			&i.ProductID,
			&i.Buyer,
			&i.Seller,
			&i.Status,
			&i.TotalAmount,
			&i.CreatedAt,
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

const getSellerTransactions = `-- name: GetSellerTransactions :many
SELECT transaction_id, product_id, buyer, seller, status, total_amount, created_at FROM transactions WHERE seller = $1 ORDER BY transaction_id LIMIT $2 OFFSET $3
`

type GetSellerTransactionsParams struct {
	Seller string `json:"seller"`
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
}

func (q *Queries) GetSellerTransactions(ctx context.Context, arg GetSellerTransactionsParams) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getSellerTransactions, arg.Seller, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transaction{}
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.TransactionID,
			&i.ProductID,
			&i.Buyer,
			&i.Seller,
			&i.Status,
			&i.TotalAmount,
			&i.CreatedAt,
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

const getTransaction = `-- name: GetTransaction :one
SELECT transaction_id, product_id, buyer, seller, status, total_amount, created_at FROM transactions WHERE transaction_id = $1 LIMIT 1
`

func (q *Queries) GetTransaction(ctx context.Context, transactionID int64) (Transaction, error) {
	row := q.db.QueryRow(ctx, getTransaction, transactionID)
	var i Transaction
	err := row.Scan(
		&i.TransactionID,
		&i.ProductID,
		&i.Buyer,
		&i.Seller,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}

const updateTransactionStatus = `-- name: UpdateTransactionStatus :one
UPDATE transactions SET status = $2 WHERE transaction_id = $1 RETURNING transaction_id, product_id, buyer, seller, status, total_amount, created_at
`

type UpdateTransactionStatusParams struct {
	TransactionID int64       `json:"transaction_id"`
	Status        pgtype.Text `json:"status"`
}

func (q *Queries) UpdateTransactionStatus(ctx context.Context, arg UpdateTransactionStatusParams) (Transaction, error) {
	row := q.db.QueryRow(ctx, updateTransactionStatus, arg.TransactionID, arg.Status)
	var i Transaction
	err := row.Scan(
		&i.TransactionID,
		&i.ProductID,
		&i.Buyer,
		&i.Seller,
		&i.Status,
		&i.TotalAmount,
		&i.CreatedAt,
	)
	return i, err
}
