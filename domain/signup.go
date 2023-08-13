package domain

import (
	"context"
	"time"
)

type SignupRequest struct {
	FullName  string    `json:"full_name" binding:"required"`
	Username  string    `json:"username" binding:"required"`
	Email     string    `json:"email" binding:"required,email"`
	Password  string    `json:"password" binding:"required"`
	BirthDate time.Time `json:"birth_date" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	// CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	// CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
