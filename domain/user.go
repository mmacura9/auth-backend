package domain

import (
	"context"
	"time"

	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
)

type User struct {
	ID        int64     `json:"id"`
	FullName  string    `json:"full_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	LastLogin time.Time `json:"last_login"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
}

func ToUserDomain(user db.User) *User {
	return &User{
		ID:        user.ID,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		BirthDate: user.BirthDate,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		LastLogin: user.LastLogin,
		DeletedAt: user.DeletedAt.Time,
	}
}
