package usecase

import (
	"context"
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
	return rtu.userRepository.GetByID(ctx, username)
}

func (rtu *refreshTokenUsecase) CreateAccessToken(user *domain.User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error) {
	accessToken, _, err = maker.CreateToken(user.Username, duration)
	return accessToken, err
}

func (rtu *refreshTokenUsecase) CreateRefreshToken(user *domain.User, duration time.Duration, maker tokenutil.Maker, c *gin.Context) (refreshToken string, err error) {
	refreshToken, payload, err := maker.CreateToken(user.Username, duration)

	if err != nil {
		return "", err
	}

	session := &domain.Session{
		ID:        payload.ID.String(),
		Username:  payload.Username,
		CreatedAt: payload.IssuedAt,
		ExpiresAt: payload.ExpiredAt,
		ClientIp:  c.ClientIP(),
		IsBlocked: false,
		UserAgent: c.Request.UserAgent(),
	}

	err = rtu.sessionRepository.Create(c, session)

	if err != nil {
		return "", err
	}
	return refreshToken, err
}
func (rtu *refreshTokenUsecase) ExtractUsernameFromToken(requestToken string) (string, error) {
	// return tokenutil.ExtractIDFromToken(requestToken, secret)
	return "", nil //TODO
}
