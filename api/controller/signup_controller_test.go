package controller

import (
	"testing"
	"time"

	mock_sqlc "github.com/ChooseCruise/choosecruise-backend/db/mock"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"go.uber.org/mock/gomock"
)

func TestSignupAPI(t *testing.T) {
	user := RandomUser()

	controller := gomock.NewController(t)
	defer controller.Finish()

	store := mock_sqlc.NewMockStore(controller)

	store.EXPECT().
		GetUserByEmail(gomock.Any(), user.Email).
		Times(1).
		Return(user, nil)

	// route.Setup(env, time.Second, store, gin)
}

func RandomUser() domain.User {
	return domain.User{
		ID:        randomutil.RandomInt(1, 1000),
		FullName:  randomutil.RandomFullName(),
		Email:     randomutil.RandomEmail(),
		Username:  randomutil.RandomUsername(),
		Password:  randomutil.RandomPassword(),
		BirthDate: randomutil.RandomBirthDate(),
		CreatedAt: time.Now(),
		LastLogin: time.Now(),
		UpdatedAt: time.Now(),
	}
}
