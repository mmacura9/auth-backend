package repository

import (
	"context"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/domain"
)

type userRepository struct {
	store db.Store
}

func NewUserRepository(store db.Store) domain.UserRepository {
	return userRepository{store: store}
}

func (urs userRepository) Create(c context.Context, user *domain.User) error {
	usr := db.CreateUserParams{
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Password:  user.Password,
		BirthDate: user.BirthDate,
	}

	_, err := urs.store.CreateUser(c, usr)
	return err
}

func (urs userRepository) GetByEmail(c context.Context, email string) (domain.User, error) {
	usr, err := urs.store.GetUserByEmail(c, email)
	if err != nil {
		return domain.User{}, err
	}
	output := domain.ToUserDomain(usr)
	return *output, err
}

func (urs userRepository) GetByUsername(c context.Context, username string) (domain.User, error) {
	usr, err := urs.store.GetUserByUsername(c, username)
	if err != nil {
		return domain.User{}, err
	}
	output := domain.ToUserDomain(usr)
	return *output, err
}
