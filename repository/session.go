package repository

import (
	"context"
	"fmt"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
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
	if err != nil {
		return domain.Session{}, err
	}

	out := domain.ToSessionDomain(session)
	return *out, err
}

func (srs sessionRepository) UpdateByID(c context.Context, maker tokenutil.Maker, id string, refreshToken string) (domain.Session, error) {
	session, err := srs.store.GetSessionByID(c, id)
	if err != nil {
		return domain.Session{}, err
	}

	payload, err := maker.VerifyToken(refreshToken)
	if err != nil {
		return domain.Session{}, err
	}

	if payload.Username != session.Username {
		return domain.Session{}, fmt.Errorf("User not authorized to extend session")
	}

	arg := db.UpdateSessionParams{
		ID:           id,
		ExpiresAt:    payload.ExpiredAt,
		RefreshToken: refreshToken,
	}
	session1, err := srs.store.UpdateSession(c, arg)
	sessionOut := domain.ToSessionDomain(session1)

	return *sessionOut, err
}
