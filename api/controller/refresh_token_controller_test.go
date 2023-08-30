package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	mock_domain "github.com/ChooseCruise/choosecruise-backend/domain/mock"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestRefreshTokenAPI(t *testing.T) {
	env := bootstrap.LoadEnv("../..")

	user := randomUser()

	testCases := []struct {
		name          string
		userEmail     string
		buildStubs    func(refreshTokenUsecase *mock_domain.MockRefreshTokenUsecase)
		checkResponce func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			userEmail: user.Email,
			buildStubs: func(refreshTokenUsecase *mock_domain.MockRefreshTokenUsecase) {
				refreshTokenUsecase.EXPECT().
					ExtractUsernameFromToken(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user.Username, nil)

				refreshTokenUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				refreshTokenUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("token1", "token2", nil)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "wrongUsername",
			userEmail: user.Email,
			buildStubs: func(refreshTokenUsecase *mock_domain.MockRefreshTokenUsecase) {
				refreshTokenUsecase.EXPECT().
					ExtractUsernameFromToken(gomock.Any(), gomock.Any()).
					Times(1).
					Return("", tokenutil.ErrorInvalidToken)

				refreshTokenUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(0)

				refreshTokenUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "NoUser",
			userEmail: user.Email,
			buildStubs: func(refreshTokenUsecase *mock_domain.MockRefreshTokenUsecase) {
				refreshTokenUsecase.EXPECT().
					ExtractUsernameFromToken(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user.Username, nil)

				refreshTokenUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				refreshTokenUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "CreateTokenError",
			userEmail: user.Email,
			buildStubs: func(refreshTokenUsecase *mock_domain.MockRefreshTokenUsecase) {
				refreshTokenUsecase.EXPECT().
					ExtractUsernameFromToken(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user.Username, nil)

				refreshTokenUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				refreshTokenUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("", "", sql.ErrConnDone)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			maker, err := tokenutil.NewPasetoMaker(env.RefreshTokenSecret)
			require.NoError(t, err)

			token, _, err := maker.CreateToken(user.Username, time.Hour)
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			refreshTokenUsecase := mock_domain.NewMockRefreshTokenUsecase(ctrl)
			testCase.buildStubs(refreshTokenUsecase)
			recorder := httptest.NewRecorder()

			url := "/api/v1/auth/refreshToken"
			values := gin.H{
				"refresh_token": token,
			}

			jsonValue, err := json.Marshal(values)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
			require.NoError(t, err)
			request.Header.Set("Content-Type", "application/json")

			rtc := RefreshTokenController{
				RefreshTokenUsecase: refreshTokenUsecase,
				Env:                 env,
				Maker:               maker,
			}
			router := gin.Default()
			router.POST(url, rtc.RefreshToken)

			router.ServeHTTP(recorder, request)
			testCase.checkResponce(t, recorder)
		})
	}
}
