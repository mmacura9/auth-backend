package repository

import (
	"context"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type SessionRepositoryStruct struct {
	store Store
}

func NewSessionRepository(store Store) domain.SessionRepository {
	return SessionRepositoryStruct{store: store}
}

func (srs SessionRepositoryStruct) Create(c context.Context, session *domain.Session) error {
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

func (srs SessionRepositoryStruct) Fetch(c context.Context) ([]domain.Session, error) {
	sessions, err := srs.store.GetAllSessions(c)
	var out []domain.Session
	for i := 0; i < len(sessions); i++ {
		session := domain.Session{
			ID:           sessions[i].ID,
			Username:     sessions[i].Username,
			RefreshToken: sessions[i].RefreshToken,
			UserAgent:    sessions[i].UserAgent,
			ClientIp:     sessions[i].ClientIp,
			IsBlocked:    sessions[i].IsBlocked,
			ExpiresAt:    sessions[i].ExpiresAt,
			CreatedAt:    sessions[i].CreatedAt,
		}
		out = append(out, session)
	}
	return out, err
}

func (srs SessionRepositoryStruct) GetByUsername(c context.Context, username string) (domain.Session, error) {
	session, err := srs.store.GetSessionByUsername(c, username)
	out := domain.Session{
		ID:           session.ID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
	return out, err
}

func (srs SessionRepositoryStruct) GetByID(c context.Context, id string) (domain.Session, error) {
	session, err := srs.store.GetSessionByID(c, id)

	out := domain.Session{
		ID:           session.ID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt,
		CreatedAt:    session.CreatedAt,
	}
	return out, err
}
