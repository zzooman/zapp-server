// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password, phone, email, location) VALUES ($1, $2, $3, $4, $5) RETURNING username, password, email, phone, location, password_changed_at, created_at
`

type CreateUserParams struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Phone    pgtype.Text `json:"phone"`
	Email    string      `json:"email"`
	Location string      `json:"location"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.Phone,
		arg.Email,
		arg.Location,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Phone,
		&i.Location,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE username = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUser, username)
	return err
}

const getUser = `-- name: GetUser :one
SELECT username, password, email, phone, location, password_changed_at, created_at FROM users WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Password,
		&i.Email,
		&i.Phone,
		&i.Location,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET password = $2, phone = $3, email = $4, location = $5 WHERE username = $1
`

type UpdateUserParams struct {
	Username string      `json:"username"`
	Password string      `json:"password"`
	Phone    pgtype.Text `json:"phone"`
	Email    string      `json:"email"`
	Location string      `json:"location"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.Exec(ctx, updateUser,
		arg.Username,
		arg.Password,
		arg.Phone,
		arg.Email,
		arg.Location,
	)
	return err
}
