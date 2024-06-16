// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: message.sql

package db

import (
	"context"
)

const createMessage = `-- name: CreateMessage :one
INSERT INTO Messages (room_id, sender, message) VALUES ($1, $2, $3) RETURNING id, room_id, sender, message, created_at, read_at
`

type CreateMessageParams struct {
	RoomID  int64  `json:"room_id"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
}

func (q *Queries) CreateMessage(ctx context.Context, arg CreateMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, createMessage, arg.RoomID, arg.Sender, arg.Message)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Sender,
		&i.Message,
		&i.CreatedAt,
		&i.ReadAt,
	)
	return i, err
}

const deleteMessage = `-- name: DeleteMessage :one
DELETE FROM Messages WHERE id = $1 RETURNING id, room_id, sender, message, created_at, read_at
`

func (q *Queries) DeleteMessage(ctx context.Context, id int64) (Message, error) {
	row := q.db.QueryRow(ctx, deleteMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Sender,
		&i.Message,
		&i.CreatedAt,
		&i.ReadAt,
	)
	return i, err
}

const getLastMessage = `-- name: GetLastMessage :one
SELECT id, room_id, sender, message, created_at, read_at FROM Messages WHERE room_id = $1 ORDER BY id DESC LIMIT 1
`

func (q *Queries) GetLastMessage(ctx context.Context, roomID int64) (Message, error) {
	row := q.db.QueryRow(ctx, getLastMessage, roomID)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Sender,
		&i.Message,
		&i.CreatedAt,
		&i.ReadAt,
	)
	return i, err
}

const getMessagesByRoom = `-- name: GetMessagesByRoom :many
SELECT id, room_id, sender, message, created_at, read_at FROM Messages WHERE room_id = $1 ORDER BY id
`

func (q *Queries) GetMessagesByRoom(ctx context.Context, roomID int64) ([]Message, error) {
	rows, err := q.db.Query(ctx, getMessagesByRoom, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Message{}
	for rows.Next() {
		var i Message
		if err := rows.Scan(
			&i.ID,
			&i.RoomID,
			&i.Sender,
			&i.Message,
			&i.CreatedAt,
			&i.ReadAt,
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

const readMessage = `-- name: ReadMessage :one
UPDATE Messages SET read_at = NOW() WHERE id = $1 RETURNING id, room_id, sender, message, created_at, read_at
`

func (q *Queries) ReadMessage(ctx context.Context, id int64) (Message, error) {
	row := q.db.QueryRow(ctx, readMessage, id)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Sender,
		&i.Message,
		&i.CreatedAt,
		&i.ReadAt,
	)
	return i, err
}

const updateMessage = `-- name: UpdateMessage :one
UPDATE Messages SET message = $2 WHERE id = $1 RETURNING id, room_id, sender, message, created_at, read_at
`

type UpdateMessageParams struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

func (q *Queries) UpdateMessage(ctx context.Context, arg UpdateMessageParams) (Message, error) {
	row := q.db.QueryRow(ctx, updateMessage, arg.ID, arg.Message)
	var i Message
	err := row.Scan(
		&i.ID,
		&i.RoomID,
		&i.Sender,
		&i.Message,
		&i.CreatedAt,
		&i.ReadAt,
	)
	return i, err
}
