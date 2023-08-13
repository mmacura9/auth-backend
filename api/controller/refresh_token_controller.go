package controller

import (
	"net/http"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
)

type RefreshTokenController struct {
	RefreshTokenUsecase domain.RefreshTokenUsecase
	Env                 *bootstrap.Env
	Maker               tokenutil.Maker
}

func (rtc *RefreshTokenController) RefreshToken(ctx *gin.Context) {
	var request domain.RefreshTokenRequest

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	username, err := rtc.RefreshTokenUsecase.ExtractUsernameFromToken(request.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, _, err := rtc.Maker.CreateToken(user.Username, time.Duration(rtc.Env.AccessTokenExpiryHour))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, _, err := rtc.Maker.CreateToken(user.Username, time.Duration(rtc.Env.RefreshTokenExpiryHour))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, refreshTokenResponse)
}
