package repository

import (
	"context"
	"testing"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) domain.Session {
	maker, err := tokenutil.NewPasetoMaker(randomutil.RandomString(32))
	require.NoError(t, err)

	Sessionname := randomutil.RandomUsername()
	duration := time.Second

	token, payload, err := maker.CreateToken(Sessionname, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	t1 := time.Now().Local().UTC()

	arg := domain.Session{
		ID:           randomutil.RandomString(10),
		Username:     randomutil.RandomUsername(),
		RefreshToken: token,
		UserAgent:    randomutil.RandomString(10),
		ClientIp:     "0.0.0.0",
		IsBlocked:    false,
		CreatedAt:    t1,
		ExpiresAt:    t1.Add(duration),
	}

	err = sessionRep.Create(context.Background(), &arg)
	require.NoError(t, err)

	err = sessionRep.Create(context.Background(), &arg)
	require.Error(t, err)

	return arg
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestGetSessionBySessionname(t *testing.T) {
	session := createRandomSession(t)

	sessions, err := sessionRep.GetByUsername(context.Background(), session.Username)
	require.NoError(t, err)
	require.NotEmpty(t, sessions)
	require.NotEmpty(t, sessions[0])
	session1 := sessions[0]

	require.Equal(t, session.ID, session1.ID)
	require.Equal(t, session.Username, session1.Username)
	require.Equal(t, session.UserAgent, session1.UserAgent)
	require.Equal(t, session.RefreshToken, session1.RefreshToken)
	require.Equal(t, session.ClientIp, session1.ClientIp)
	require.Equal(t, session.IsBlocked, session1.IsBlocked)
	require.WithinDuration(t, session.CreatedAt, session1.CreatedAt, time.Second)
	require.WithinDuration(t, session.ExpiresAt, session1.ExpiresAt, time.Second)
}
