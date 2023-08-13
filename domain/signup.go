package domain

import (
	"context"
	"time"
)

type SignupRequest struct {
	Full_name  string    `form:"full_name" binding:"required"`
	Username   string    `form:"full_name" binding:"required"`
	Email      string    `form:"email" binding:"required,email"`
	Password   string    `form:"password" binding:"required"`
	Birth_date time.Time `form:"birth_date" binding:"required"`
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
