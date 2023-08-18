package db

import (
	"context"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:  randomutil.RandomUsername(),
		FullName:  randomutil.RandomFullName(),
		Password:  randomutil.RandomPassword(),
		Email:     randomutil.RandomEmail(),
		BirthDate: randomutil.RandomBirthDate(),
	}

	user1, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.NotZero(t, user1.ID)
	require.NotZero(t, user1.CreatedAt)
	require.NotZero(t, user1.LastLogin)
	require.NotZero(t, user1.UpdatedAt)

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, user)
	return user1
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testQueries.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.NotZero(t, user1.ID)
	require.NotZero(t, user1.Username)
	require.NotZero(t, user1.FullName)
	require.NotZero(t, user1.Email)
	require.NotZero(t, user1.Password)
	require.NotZero(t, user1.CreatedAt)
	require.NotZero(t, user1.LastLogin)
	require.NotZero(t, user1.UpdatedAt)

	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Password, user1.Password)

}
