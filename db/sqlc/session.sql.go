// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: session.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createSession = `-- name: CreateSession :execresult
INSERT INTO sessions (
    id,
    username,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at
) VALUES (
$1, $2, $3, $4, $5, $6, $7
)
`

type CreateSessionParams struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createSession,
		arg.ID,
		arg.Username,
		arg.RefreshToken,
		arg.UserAgent,
		arg.ClientIp,
		arg.IsBlocked,
		arg.ExpiresAt,
	)
}

const getAllSessions = `-- name: GetAllSessions :many
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
FROM sessions
`

func (q *Queries) GetAllSessions(ctx context.Context) ([]Session, error) {
	rows, err := q.db.QueryContext(ctx, getAllSessions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Session{}
	for rows.Next() {
		var i Session
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.RefreshToken,
			&i.UserAgent,
			&i.ClientIp,
			&i.IsBlocked,
			&i.ExpiresAt,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getSessionByID = `-- name: GetSessionByID :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
FROM sessions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetSessionByID(ctx context.Context, id string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByID, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSessionByUsername = `-- name: GetSessionByUsername :one
SELECT id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at, created_at
FROM sessions
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetSessionByUsername(ctx context.Context, username string) (Session, error) {
	row := q.db.QueryRowContext(ctx, getSessionByUsername, username)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.RefreshToken,
		&i.UserAgent,
		&i.ClientIp,
		&i.IsBlocked,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}