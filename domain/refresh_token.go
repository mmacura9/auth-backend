package domain

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenUsecase interface {
	GetUserByUsername(c context.Context, id string) (User, error)
	CreateTokens(c *gin.Context, user *User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker, oldRefreshToken string) (accessToken string, refreshToken string, err error)
	ExtractUsernameFromToken(requestToken string, maker tokenutil.Maker) (string, error)
}
