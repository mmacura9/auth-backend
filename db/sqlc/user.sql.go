// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO "user" ("username", "email", "full_name", "password", "birth_date")
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
`

type CreateUserParams struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.FullName,
		arg.Password,
		arg.BirthDate,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.BirthDate,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.DeletedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "user"
WHERE "username" = $1
`

func (q *Queries) DeleteUser(ctx context.Context, username string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, username)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
FROM "user"
WHERE "email" = $1 LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.BirthDate,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.DeletedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
FROM "user"
WHERE "username" = $1 LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.BirthDate,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.DeletedAt,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
FROM "user"
WHERE "username" = $1 LIMIT 1 FOR UPDATE
`

func (q *Queries) GetUserForUpdate(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserForUpdate, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.BirthDate,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.DeletedAt,
	)
	return i, err
}

const listUser = `-- name: ListUser :many
SELECT id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
FROM "user"
ORDER BY "username"
`

func (q *Queries) ListUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.FullName,
			&i.BirthDate,
			&i.Password,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.LastLogin,
			&i.DeletedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE "user" 
SET "username"=$1, "email"=$2, "full_name"=$3, "password"=$4
WHERE "username" = $5
RETURNING id, username, email, full_name, birth_date, password, created_at, updated_at, last_login, deleted_at
`

type UpdateUserParams struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Password   string `json:"password"`
	Username_2 string `json:"username_2"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Username,
		arg.Email,
		arg.FullName,
		arg.Password,
		arg.Username_2,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.BirthDate,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastLogin,
		&i.DeletedAt,
	)
	return i, err
}
