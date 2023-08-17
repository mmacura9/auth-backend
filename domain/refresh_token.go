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
	// createAccessToken(user *User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error)
	// createRefreshToken(user *User, duration time.Duration, maker tokenutil.Maker, c *gin.Context) (refreshToken string, err error)
	CreateTokens(user *User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker, c *gin.Context) (accessToken string, refreshToken string, err error)
	ExtractUsernameFromToken(requestToken string, maker tokenutil.Maker) (string, error)
}
