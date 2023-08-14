package domain

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refreshToken" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenUsecase interface {
	GetUserByUsername(c context.Context, id string) (User, error)
	CreateAccessToken(user *User, duration time.Duration, maker tokenutil.Maker) (refreshToken string, err error)
	CreateRefreshToken(user *User, duration time.Duration, maker tokenutil.Maker, c *gin.Context) (refreshToken string, err error)
	ExtractUsernameFromToken(requestToken string) (string, error)
}
