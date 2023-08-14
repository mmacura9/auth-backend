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

func (lu *loginUsecase) CreateAccessToken(user *domain.User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error) {
	accessToken, _, err = maker.CreateToken(user.Username, duration)
	return accessToken, err
}

func (lu *loginUsecase) CreateRefreshToken(user *domain.User, duration time.Duration, maker tokenutil.Maker, c *gin.Context) (refreshToken string, err error) {
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

	err = lu.sessionRepository.Create(c, session)

	if err != nil {
		return "", err
	}
	return refreshToken, err
}
