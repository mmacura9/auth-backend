package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	"github.com/ChooseCruise/choosecruise-backend/domain"
	mock_domain "github.com/ChooseCruise/choosecruise-backend/domain/mock"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

func TestLoginAPI(t *testing.T) {
	env := bootstrap.LoadEnv("../..")

	user := randomUser()

	testCases := []struct {
		name          string
		userEmail     string
		buildStubs    func(loginUsecase *mock_domain.MockLoginUsecase)
		checkResponce func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			userEmail: user.Email,
			buildStubs: func(loginUsecase *mock_domain.MockLoginUsecase) {
				encryptedPassword, err := bcrypt.GenerateFromPassword(
					[]byte(user.Password),
					bcrypt.DefaultCost,
				)
				require.NoError(t, err)

				user1 := user
				user1.Password = string(encryptedPassword)
				loginUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user1, nil)

				loginUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("token1", "token2", nil)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "UserNotFound",
			userEmail: user.Email,
			buildStubs: func(loginUsecase *mock_domain.MockLoginUsecase) {
				loginUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				loginUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
		{
			name:      "CreateTokensError",
			userEmail: user.Email,
			buildStubs: func(loginUsecase *mock_domain.MockLoginUsecase) {
				encryptedPassword, err := bcrypt.GenerateFromPassword(
					[]byte(user.Password),
					bcrypt.DefaultCost,
				)
				require.NoError(t, err)

				user1 := user
				user1.Password = string(encryptedPassword)
				loginUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user1, nil)

				loginUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("", "", sql.ErrConnDone)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "BadRequest",
			userEmail: "invalid-email",
			buildStubs: func(loginUsecase *mock_domain.MockLoginUsecase) {
				loginUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(0)

				loginUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		{
			name:      "WrongPassword",
			userEmail: user.Email,
			buildStubs: func(loginUsecase *mock_domain.MockLoginUsecase) {
				loginUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				loginUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			maker, err := tokenutil.NewPasetoMaker(env.RefreshTokenSecret)
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			loginUsecase := mock_domain.NewMockLoginUsecase(ctrl)
			testCase.buildStubs(loginUsecase)
			recorder := httptest.NewRecorder()

			url := "/api/v1/auth/login"
			values := gin.H{
				"email":    testCase.userEmail,
				"password": user.Password,
			}

			jsonValue, err := json.Marshal(values)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
			require.NoError(t, err)

			request.Header.Set("Content-Type", "application/json")

			lc := LoginController{
				LoginUsecase: loginUsecase,
				Env:          env,
				Maker:        maker,
			}
			router := gin.Default()
			router.POST(url, lc.Login)

			router.ServeHTTP(recorder, request)
			testCase.checkResponce(t, recorder)
		})
	}
}
