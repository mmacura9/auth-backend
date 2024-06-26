package usecase

import (
	"context"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type signupUsecase struct {
	userRepository    domain.UserRepository
	sessionRepository domain.SessionRepository
	contextTimeout    time.Duration
}

func NewSignupUsecase(userRepository domain.UserRepository, sessionRepository domain.SessionRepository, timeout time.Duration) domain.SignupUsecase {
	return &signupUsecase{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
		contextTimeout:    timeout,
	}
}

func (su *signupUsecase) Create(c context.Context, user *domain.User) error {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.Create(ctx, user)
}

func (su *signupUsecase) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByEmail(ctx, email)
}

func (su *signupUsecase) GetUserByUsername(c context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, su.contextTimeout)
	defer cancel()
	return su.userRepository.GetByUsername(ctx, email)
}

func (su *signupUsecase) CreateTokens(c *gin.Context, user *domain.User, accDuration time.Duration, refDuration time.Duration, maker tokenutil.Maker) (accessToken string, refreshToken string, err error) {
	accessToken, err = su.createAccessToken(user, accDuration, maker)
	if err != nil {
		return "", "", err
	}

	refreshToken, err = su.createRefreshToken(c, user, refDuration, maker)
	return
}

func (su *signupUsecase) createAccessToken(user *domain.User, duration time.Duration, maker tokenutil.Maker) (accessToken string, err error) {
	accessToken, _, err = maker.CreateToken(user.Username, duration)
	return accessToken, err
}

func (su *signupUsecase) createRefreshToken(c *gin.Context, user *domain.User, duration time.Duration, maker tokenutil.Maker) (refreshToken string, err error) {
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

	err = su.sessionRepository.Create(c, session)

	if err != nil {
		return "", err
	}
	return refreshToken, err
}
