package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{Queries: New(db), db: db}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rberr := tx.Rollback(); rberr != nil {
			return fmt.Errorf("tx err %v, rollback error %v", err, rberr)
		}
		return err
	}

	return tx.Commit()
}

type UserRepositoryStruct struct {
	db Store
}

func NewUserRepository(db Store) domain.UserRepository {
	return UserRepositoryStruct{db: db}
}

func (urs UserRepositoryStruct) Create(c context.Context, user *domain.User) error {
	usr := CreateUserParams{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Password:  user.Password,
		BirthDate: user.BirthDate,
	}
	fmt.Println(usr)
	_, err := urs.db.CreateUser(c, usr)
	return err
}

func (urs UserRepositoryStruct) Fetch(c context.Context) ([]domain.User, error) {

	return nil, nil
}

func (urs UserRepositoryStruct) GetByEmail(c context.Context, email string) (domain.User, error) {
	usr, err := urs.db.GetUserByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	output := domain.User{
		ID:        usr.ID,
		Username:  usr.Username,
		Password:  usr.Password,
		Email:     usr.Email,
		FullName:  usr.FullName,
		BirthDate: usr.BirthDate,
		CreatedAt: usr.CreatedAt,
		UpdatedAt: usr.UpdatedAt,
		LastLogin: usr.LastLogin,
	}
	return output, err
}

func (urs UserRepositoryStruct) GetByID(c context.Context, id string) (domain.User, error) {
	return domain.User{}, nil
}

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
