package domain

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginUsecase interface {
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateTokens(c *gin.Context, user *User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker) (accessToken string, refreshToken string, err error)
}
