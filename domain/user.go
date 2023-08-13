package domain

import (
	"context"
	"time"
)

type User struct {
	ID         int64     `json:"_id"`
	Full_name  string    `json:"full_name"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Birth_date time.Time `json:"birth_date"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
	Last_login time.Time `json:"last_login"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
