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
		ExpiresAt:    time.Now().Add(duration),
	}

	session, err := testQueries.CreateSession(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	session1, err := testQueries.CreateSession(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, session1)
	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}
