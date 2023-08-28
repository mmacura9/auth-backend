package domain

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
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
	GetUserByUsername(c context.Context, username string) (User, error)
	CreateTokens(c *gin.Context, user *User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker) (accessToken string, refreshToken string, err error)
}
