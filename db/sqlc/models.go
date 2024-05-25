// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID                int64              `json:"id"`
	Owner             string             `json:"owner"`
	AccountNumber     string             `json:"account_number"`
	BankName          string             `json:"bank_name"`
	AccountHolderName string             `json:"account_holder_name"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

type Comment struct {
	ID              int64              `json:"id"`
	PostID          int64              `json:"post_id"`
	ParentCommentID pgtype.Int8        `json:"parent_comment_id"`
	Commentor       string             `json:"commentor"`
	CommentText     string             `json:"comment_text"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
}

type LikeWithPost struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}

type Payment struct {
	PaymentID     int64              `json:"payment_id"`
	TransactionID int64              `json:"transaction_id"`
	PaymentStatus pgtype.Text        `json:"payment_status"`
	PaymentMethod string             `json:"payment_method"`
	PaymentDate   pgtype.Timestamptz `json:"payment_date"`
	PaymentAmount pgtype.Numeric     `json:"payment_amount"`
}

type Post struct {
	ID        int64              `json:"id"`
	Author    string             `json:"author"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Media     []string           `json:"media"`
	Price     int64              `json:"price"`
	Stock     int64              `json:"stock"`
	Views     pgtype.Int8        `json:"views"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Review struct {
	ID        int64              `json:"id"`
	Seller    string             `json:"seller"`
	Reviewer  string             `json:"reviewer"`
	Rating    int32              `json:"rating"`
	Content   pgtype.Text        `json:"content"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Transaction struct {
	TransactionID int64              `json:"transaction_id"`
	PostID        int64              `json:"post_id"`
	Buyer         string             `json:"buyer"`
	Seller        string             `json:"seller"`
	Status        pgtype.Text        `json:"status"`
	TotalAmount   pgtype.Numeric     `json:"total_amount"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	Username          string             `json:"username"`
	Password          string             `json:"password"`
	Email             string             `json:"email"`
	Phone             pgtype.Text        `json:"phone"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}

type WishWithProduct struct {
	Username string `json:"username"`
	PostID   int64  `json:"post_id"`
}
