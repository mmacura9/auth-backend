package domain

import (
	"context"
	"time"
)

type Session struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type SessionRepository interface {
	Create(c context.Context, session *Session) error
	Fetch(c context.Context) ([]Session, error)
	GetByUsername(c context.Context, email string) ([]Session, error)
	GetByID(c context.Context, id string) (Session, error)
}
