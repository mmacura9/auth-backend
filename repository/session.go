package repository

import (
	"context"

	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type SessionRepositoryStruct struct {
	db Store
}

func NewSessionRepository(db Store) domain.SessionRepository {
	return SessionRepositoryStruct{db: db}
}

func (srs SessionRepositoryStruct) Create(c context.Context, session *domain.Session) error {
	params := CreateSessionParams{
		ID:           session.ID,
		Username:     session.Username,
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    session.ExpiresAt,
	}

	_, err := srs.db.CreateSession(c, params)

	return err
}

func (srs SessionRepositoryStruct) Fetch(c context.Context) ([]domain.Session, error) {
	sessions, err := srs.db.GetAllSessions(c)
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
	session, err := srs.db.GetSessionByUsername(c, username)
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
	session, err := srs.db.GetSessionByID(c, id)

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
