package repository

import (
	"context"
	"fmt"

	"github.com/ChooseCruise/choosecruise-backend/domain"
)

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
