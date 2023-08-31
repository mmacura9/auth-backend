// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"
)

type Querier interface {
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteUser(ctx context.Context, username string) error
	GetSessionByID(ctx context.Context, id string) (Session, error)
	GetSessionByUsername(ctx context.Context, username string) ([]Session, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	UpdateSession(ctx context.Context, arg UpdateSessionParams) (Session, error)
}

var _ Querier = (*Queries)(nil)
