package controller

import (
	"net/http"

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
		ctx.JSON(http.StatusBadRequest, domain.NewErrorResponse(err.Error()))
		return
	}

	username, err := rtc.RefreshTokenUsecase.ExtractUsernameFromToken(request.RefreshToken, rtc.Maker)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("User not found"))
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByUsername(ctx, username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, domain.NewErrorResponse("User not found"))
		return
	}

	accessToken, refreshToken, err := rtc.RefreshTokenUsecase.CreateTokens(ctx, &user, rtc.Env.AccessTokenExpiry, rtc.Env.RefreshTokenExpiry, rtc.Maker, request.RefreshToken)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	ctx.JSON(http.StatusOK, refreshTokenResponse)
}
