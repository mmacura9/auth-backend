package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type refreshTokenUsecase struct {
	userRepository    domain.UserRepository
	sessionRepository domain.SessionRepository
	contextTimeout    time.Duration
}

func NewRefreshTokenUsecase(userRepository domain.UserRepository, sessionRepository domain.SessionRepository, timeout time.Duration) domain.RefreshTokenUsecase {
	return &refreshTokenUsecase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		contextTimeout:    timeout,
	}
}

func (rtu *refreshTokenUsecase) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, rtu.contextTimeout)
	defer cancel()
	return rtu.userRepository.GetByUsername(ctx, username)
}

func (rtu *refreshTokenUsecase) ExtractUsernameFromToken(requestToken string, maker tokenutil.Maker) (string, error) {
	payload, err := maker.VerifyToken(requestToken)
	if err != nil {
		return "", err
	}

	return payload.Username, nil
}

func (rtu *refreshTokenUsecase) CreateTokens(c *gin.Context, user *domain.User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker, oldRefreshToken string) (accessToken string, refreshToken string, err error) {
	accessToken, err = rtu.createAccessToken(user, accDuration, maker)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = rtu.createRefreshToken(c, user, refDuration, maker, oldRefreshToken)
	return
}

func (rtu *refreshTokenUsecase) createAccessToken(user *domain.User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error) {
	accessToken, _, err = maker.CreateToken(user.Username, duration)
	return accessToken, err
}

func (rtu *refreshTokenUsecase) createRefreshToken(c *gin.Context, user *domain.User, duration time.Duration, maker tokenutil.Maker, oldRefreshToken string) (string, error) {
	refreshToken, _, err := maker.CreateToken(user.Username, duration)
	if err != nil {
		return "", err
	}

	sessions, err := rtu.sessionRepository.GetByUsername(c, user.Username)
	if err != nil {
		return "", err
	}

	var session *domain.Session

	for _, sessionEntity := range sessions {
		if sessionEntity.RefreshToken == oldRefreshToken {
			session = &sessionEntity
			break
		}
	}

	if session == nil {
		// TODO: Delete all refresh tokens for this user
		return "", fmt.Errorf("Old refresh token!")
	}

	_, err = rtu.sessionRepository.UpdateByID(c, maker, session.ID, refreshToken)
	if err != nil {
		return "", err
	}

	return refreshToken, err
}
