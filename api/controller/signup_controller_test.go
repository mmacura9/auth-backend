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
	"github.com/ChooseCruise/choosecruise-backend/internal/randomutil"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestSignupAPI(t *testing.T) {
	env := bootstrap.LoadEnv("../..")

	user := randomUser()

	testCases := []struct {
		name          string
		userEmail     string
		buildStubs    func(signupUsecase *mock_domain.MockSignupUsecase)
		checkResponce func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:      "OK",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("token1", "token2", nil)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{
			name:      "FoundEmail",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)
				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name:      "FoundUsername",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(user, nil)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusConflict, recorder.Code)
			},
		},
		{
			name:      "CreateError",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "CreateTokensError",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return("", "", sql.ErrConnDone)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:      "CreateTokensError",
			userEmail: user.Email,
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(1).
					Return(domain.User{}, sql.ErrNoRows)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(1).
					Return(nil)

				signupUsecase.EXPECT().
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
			buildStubs: func(signupUsecase *mock_domain.MockSignupUsecase) {
				signupUsecase.EXPECT().
					GetUserByEmail(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					GetUserByUsername(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					Create(gomock.Any(), gomock.Any()).
					Times(0)

				signupUsecase.EXPECT().
					CreateTokens(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponce: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			maker, err := tokenutil.NewPasetoMaker(env.RefreshTokenSecret)
			require.NoError(t, err)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			signupUsecase := mock_domain.NewMockSignupUsecase(ctrl)
			testCase.buildStubs(signupUsecase)
			recorder := httptest.NewRecorder()

			url := "/api/v1/auth/signup"
			values := gin.H{
				"full_name":  user.FullName,
				"username":   user.Username,
				"email":      testCase.userEmail,
				"password":   user.Password,
				"birth_date": user.BirthDate,
			}

			jsonValue, err := json.Marshal(values)
			require.NoError(t, err)

			request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonValue))
			require.NoError(t, err)

			sc := SignupController{
				SignupUsecase: signupUsecase,
				Env:           env,
				Maker:         maker,
			}
			router := gin.Default()
			router.POST(url, sc.Signup)

			router.ServeHTTP(recorder, request)
			testCase.checkResponce(t, recorder)
		})
	}
}

func randomUser() domain.User {
	return domain.User{
		ID:        randomutil.RandomInt(1, 1000),
		FullName:  randomutil.RandomFullName(),
		Username:  randomutil.RandomUsername(),
		Email:     randomutil.RandomEmail(),
		Password:  randomutil.RandomPassword(),
		BirthDate: randomutil.RandomBirthDate(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		LastLogin: time.Now(),
	}
}
