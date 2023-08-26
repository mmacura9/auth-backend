package repository

import (
	"context"
	"testing"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) domain.User {
	arg := domain.User{
		Username:  randomutil.RandomUsername(),
		FullName:  randomutil.RandomFullName(),
		Password:  randomutil.RandomPassword(),
		Email:     randomutil.RandomEmail(),
		BirthDate: randomutil.RandomBirthDate(),
	}

	err := userRep.Create(context.Background(), &arg)
	require.NoError(t, err)

	err = userRep.Create(context.Background(), &arg)
	require.Error(t, err)
	return arg
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := createRandomUser(t)

	user1, err := userRep.GetByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user1)

	require.Equal(t, user.Email, user1.Email)
	require.Equal(t, user.Username, user1.Username)
	require.Equal(t, user.FullName, user1.FullName)
	require.Equal(t, user.Password, user1.Password)
	require.WithinDuration(t, user.BirthDate, user1.BirthDate, time.Second)

}
