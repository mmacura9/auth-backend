package db

import (
	"context"
	"testing"
	"time"

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

	user1, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.NotZero(t, user1.ID)
	require.NotZero(t, user1.CreatedAt)
	require.NotZero(t, user1.LastLogin)
	require.NotZero(t, user1.UpdatedAt)

	user, err := testStore.CreateUser(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, user)
	return user1
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testStore.GetUserByEmail(context.Background(), user.Email)
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
	require.WithinDuration(t, user.BirthDate, user1.BirthDate, time.Second)
	require.WithinDuration(t, user.CreatedAt, user1.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, user1.UpdatedAt, time.Second)
	require.WithinDuration(t, user.LastLogin, user1.LastLogin, time.Second)
}

func TestDeleteUser(t *testing.T) {
	user := createRandomUser(t)

	err := testStore.DeleteUser(context.Background(), user.Username)
	require.NoError(t, err)

	user1, err := testStore.GetUserByEmail(context.Background(), user.Email)
	require.Error(t, err)
	require.Empty(t, user1)
}

func TestGetUserByUsername(t *testing.T) {
	user := createRandomUser(t)

	user1, err := testStore.GetUserByUsername(context.Background(), user.Username)
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
	require.WithinDuration(t, user.BirthDate, user1.BirthDate, time.Second)
	require.WithinDuration(t, user.CreatedAt, user1.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, user1.UpdatedAt, time.Second)
	require.WithinDuration(t, user.LastLogin, user1.LastLogin, time.Second)
}
