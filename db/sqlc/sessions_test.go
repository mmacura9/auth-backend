package db

import (
	"context"
	"testing"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) Session {
	maker, err := tokenutil.NewPasetoMaker(randomutil.RandomString(32))
	require.NoError(t, err)

	username := randomutil.RandomUsername()
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	arg := CreateSessionParams{
		ID:           randomutil.RandomUsername(),
		Username:     randomutil.RandomUsername(),
		RefreshToken: token,
		UserAgent:    randomutil.RandomString(10),
		ClientIp:     "0.0.0.0",
		IsBlocked:    false,
		ExpiresAt:    time.Now().UTC().Add(duration),
	}

	session, err := testStore.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	session1, err := testStore.CreateSession(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, session1)
	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestGetSessionById(t *testing.T) {
	session := createRandomSession(t)

	s1, err := testStore.GetSessionByID(context.Background(), session.ID)
	require.NoError(t, err)
	require.NotEmpty(t, s1)

	require.Equal(t, session.ID, s1.ID)
	require.Equal(t, session.Username, s1.Username)
	require.Equal(t, session.UserAgent, s1.UserAgent)
	require.Equal(t, session.RefreshToken, s1.RefreshToken)
	require.Equal(t, session.ClientIp, s1.ClientIp)
	require.Equal(t, session.IsBlocked, s1.IsBlocked)
	require.WithinDuration(t, session.CreatedAt, s1.CreatedAt, time.Second)
	require.WithinDuration(t, session.ExpiresAt, s1.ExpiresAt, time.Second)
}

func TestGetSessionByUsername(t *testing.T) {
	session := createRandomSession(t)

	session1, err := testStore.GetSessionByUsername(context.Background(), session.Username)
	require.NoError(t, err)
	require.NotEmpty(t, session1)

	s1 := session1[0]
	require.NotEmpty(t, s1)
	require.Equal(t, session.ID, s1.ID)
	require.Equal(t, session.Username, s1.Username)
	require.Equal(t, session.UserAgent, s1.UserAgent)
	require.Equal(t, session.RefreshToken, s1.RefreshToken)
	require.Equal(t, session.ClientIp, s1.ClientIp)
	require.Equal(t, session.IsBlocked, s1.IsBlocked)
	require.WithinDuration(t, session.CreatedAt, s1.CreatedAt, time.Second)
	require.WithinDuration(t, session.ExpiresAt, s1.ExpiresAt, time.Second)
}

func TestUpdateSession(t *testing.T) {
	session := createRandomSession(t)
	maker, err := tokenutil.NewPasetoMaker(randomutil.RandomString(32))
	require.NoError(t, err)

	username := randomutil.RandomUsername()
	duration := time.Minute

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	arg := UpdateSessionParams{
		ID:           session.ID,
		RefreshToken: token,
		ExpiresAt:    time.Now().UTC().Add(duration),
	}

	session1, err := testQueries.UpdateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session1)
}
