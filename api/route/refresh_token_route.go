package route

import (
	"time"

	"github.com/ChooseCruise/choosecruise-backend/api/controller"
	"github.com/ChooseCruise/choosecruise-backend/bootstrap"
	db "github.com/ChooseCruise/choosecruise-backend/db/sqlc"
	"github.com/ChooseCruise/choosecruise-backend/internal/tokenutil"
	"github.com/ChooseCruise/choosecruise-backend/repository"
	"github.com/ChooseCruise/choosecruise-backend/usecase"
	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, store db.Store, group *gin.RouterGroup, maker tokenutil.Maker) {
	ur := repository.NewUserRepository(store)
	sr := repository.NewSessionRepository(store)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, sr, timeout),
		Env:                 env,
		Maker:               maker,
	}
	group.POST("/refreshToken", rtc.RefreshToken)
}
