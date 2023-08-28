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

func (rtc *RefreshTokenController) RefreshToken(c *gin.Context) {
	var request domain.RefreshTokenRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.NewErrorResponse(err.Error()))
		return
	}

	username, err := rtc.RefreshTokenUsecase.ExtractUsernameFromToken(request.RefreshToken, rtc.Maker)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("User not found"))
		return
	}

	user, err := rtc.RefreshTokenUsecase.GetUserByUsername(c, username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.NewErrorResponse("User not found"))
		return
	}

	accessToken, refreshToken, err := rtc.RefreshTokenUsecase.CreateTokens(c, &user, rtc.Env.AccessTokenExpiry, rtc.Env.RefreshTokenExpiry, rtc.Maker, request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	refreshTokenResponse := domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, refreshTokenResponse)
}
