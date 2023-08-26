package repository

import (
	"context"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type sessionRepository struct {
	store db.Store
}

func NewSessionRepository(store db.Store) domain.SessionRepository {
	return sessionRepository{store: store}
}

func (srs sessionRepository) Create(c context.Context, session *domain.Session) error {
	params := db.CreateSessionParams{
		ID:           session.ID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt,
	}

	_, err := srs.store.CreateSession(c, params)

	return err
}

func (srs sessionRepository) Fetch(c context.Context) ([]domain.Session, error) {
	sessions, err := srs.store.GetAllSessions(c)
	var out []domain.Session
	for _, sessionEntity := range sessions {
		session := domain.ToSessionDomain(sessionEntity)
		out = append(out, *session)
	}
	return out, err
}

func (srs sessionRepository) GetByUsername(c context.Context, username string) ([]domain.Session, error) {
	sessions, err := srs.store.GetSessionByUsername(c, username)
	var out []domain.Session

	for _, sessionEntity := range sessions {
		session := domain.ToSessionDomain(sessionEntity)
		out = append(out, *session)
	}

	return out, err
}

func (srs sessionRepository) GetByID(c context.Context, id string) (domain.Session, error) {
	session, err := srs.store.GetSessionByID(c, id)

	out := domain.ToSessionDomain(session)
	return *out, err
}
