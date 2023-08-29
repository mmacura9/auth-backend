package usecase

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type loginUsecase struct {
	userRepository    domain.UserRepository
	sessionRepository domain.SessionRepository
	contextTimeout    time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository, sessionRepository domain.SessionRepository, timeout time.Duration) domain.LoginUsecase {
	return &loginUsecase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		contextTimeout:    timeout,
	}
}

func (lu *loginUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, lu.contextTimeout)
	defer cancel()
	return lu.userRepository.GetByEmail(ctx, email)
}

func (lu *loginUsecase) CreateTokens(c *gin.Context, user *domain.User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker) (accessToken string, refreshToken string, err error) {
	accessToken, err = lu.createAccessToken(user, accDuration, maker)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = lu.createRefreshToken(user, refDuration, maker, c)
	return
}

func (lu *loginUsecase) createAccessToken(user *domain.User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error) {
	accessToken, _, err = maker.CreateToken(user.Username, duration)
	return accessToken, err
}

func (lu *loginUsecase) createRefreshToken(user *domain.User, duration time.Duration, maker tokenutil.Maker, c *gin.Context) (refreshToken string, err error) {
	refreshToken, payload, err := maker.CreateToken(user.Username, duration)

	if err != nil {
		return "", err
	}

	session := &domain.Session{
		ID:           payload.ID.String(),
		RefreshToken: refreshToken,
		Username:     payload.Username,
		CreatedAt:    payload.IssuedAt,
		ExpiresAt:    payload.ExpiredAt,
		ClientIp:     c.ClientIP(),
		IsBlocked:    false,
		UserAgent:    c.Request.UserAgent(),
	}

	err = lu.sessionRepository.Create(c, session)

	if err != nil {
		return "", err
	}
	return refreshToken, err
}
