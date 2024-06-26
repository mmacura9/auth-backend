package domain

import (
	"context"
	"time"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
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
	GetByUsername(c context.Context, username string) ([]Session, error)
	GetByID(c context.Context, id string) (Session, error)
	UpdateByID(c context.Context, maker tokenutil.Maker, id string, refreshToken string) (Session, error)
}

func ToSessionDomain(session db.Session) *Session {
	return &Session{
		ID:           session.ID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
}
