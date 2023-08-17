package db

import (
	"context"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Username:  randomutil.RandomUsername(),
		FullName:  randomutil.RandomFullName(),
		Password:  randomutil.RandomPassword(),
		Email:     randomutil.RandomEmail(),
		BirthDate: randomutil.RandomBirthDate(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.LastLogin)
	require.NotZero(t, user.UpdatedAt)

	user, err = testQueries.CreateUser(context.Background(), arg)
	require.Error(t, err)
	require.Empty(t, user)

}
