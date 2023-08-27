package controller

import (
	"net/http"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
	Maker         tokenutil.Maker
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.NewErrorResponse(err.Error()))
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewErrorResponse("User already exists with the given email"))
		return
	}

	_, err = sc.SignupUsecase.GetUserByUsername(c, request.Username)
	if err == nil {
		c.JSON(http.StatusConflict, domain.NewErrorResponse("User already exists with the given username"))
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		FullName:  request.FullName,
		Username:  request.Username,
		Email:     request.Email,
		Password:  request.Password,
		BirthDate: request.BirthDate,
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	accessToken, refreshToken, err := sc.SignupUsecase.CreateTokens(&user, sc.Env.AccessTokenExpiry, sc.Env.RefreshTokenExpiry, sc.Maker, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.NewErrorResponse(err.Error()))
		return
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, signupResponse)
}
